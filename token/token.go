package token

import (
	"context"
	"github.com/blueambition/node-lib/account"
	token2 "github.com/blueambition/node-lib/contract/token"
	"github.com/blueambition/node-lib/custom"
	"github.com/blueambition/node-lib/transaction"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	dc "github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
	"strings"
)

//代币余额
func BalanceOf(client *ethclient.Client, contract string, address string) (float64, error) {
	token, err := token2.NewStandardToken(common.HexToAddress(contract), client)
	if err != nil {
		return 0, errors.New("获取代币信息失败：" + err.Error())
	}
	balance, err := token.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return 0, errors.New("获取余额失败：" + err.Error())
	}
	decimal, err := token.Decimals(nil)
	if err != nil {
		return 0, errors.New("获取精度失败：" + err.Error())
	}
	amountDC, _ := dc.NewFromString(balance.String())
	ratioDC := dc.NewFromFloat(math.Pow(10, float64(decimal)))
	realAmount, _ := amountDC.Div(ratioDC).RoundFloor(6).Float64()
	return realAmount, nil
}

//代币Big余额
func BigBalanceOf(client *ethclient.Client, contract string, address string) *big.Int {
	token, _ := token2.NewStandardToken(common.HexToAddress(contract), client)
	balance, err := token.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		return nil
	}
	return balance
}

//发送代币
func SendToken(client *ethclient.Client, authData custom.AuthData, contract, to string, amount float64) (*types.Transaction, error) {
	var (
		err error
		tx  *types.Transaction
	)
	if strings.EqualFold(authData.Account, to) {
		return nil, errors.New("付款地址和收款地址一致")
	}
	balance, err := BalanceOf(client, contract, authData.Account)
	if err != nil {
		return nil, err
	}
	if amount > balance {
		return nil, errors.New("余额不足")
	}
	if !account.IsErc20Address(to) {
		return nil, errors.New("非法收款地址")
	}
	token, _ := token2.NewStandardToken(common.HexToAddress(contract), client)
	decimal, _ := token.Decimals(nil)
	amountDC := dc.NewFromFloat(amount)
	ratioDC := dc.NewFromFloat(math.Pow(10, float64(decimal)))
	transAmount := amountDC.Mul(ratioDC).BigInt()
	toHex := common.HexToAddress(to)
	authOpts, err := account.AuthOpts(client, authData)
	if err != nil {
		return nil, err
	}
	tx, err = token.Transfer(authOpts, toHex, transAmount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//授权
func Approve(client *ethclient.Client, authData custom.AuthData, fromToken, contract string, amount float64) (*types.Transaction, error) {
	token, _ := token2.NewStandardToken(common.HexToAddress(fromToken), client)
	decimal, _ := token.Decimals(nil)
	amountDC := dc.NewFromFloat(amount)
	baseDC := dc.NewFromFloat(math.Pow(10, float64(decimal)))
	amountDC = amountDC.Mul(baseDC)
	amountBig := amountDC.BigInt()
	authOpts, err := account.AuthOpts(client, authData)
	if err != nil {
		return nil, err
	}
	txInfo, err := token.Approve(authOpts, common.HexToAddress(contract), amountBig)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

//交易信息
func TransInfo(client *ethclient.Client, txHash string) (custom.TransTokenInfo, error) {
	var info custom.TransTokenInfo
	realHash := common.HexToHash(txHash)
	abiJson := token2.StandardTokenABI
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return info, err
	}
	trans, isPending, err := client.TransactionByHash(context.Background(), realHash)
	if err != nil {
		return info, err
	}
	if isPending {
		return info, errors.New("交易待处理，请稍后")
	}
	//归集账户
	transTo := trans.To()
	if transTo == nil {
		return info, errors.New("交易合约地址不存在")
	}
	info.Contract = trans.To().String()
	//解析
	inputData := hexutil.Bytes(trans.Data())
	inputStr := inputData.String()
	unpackData, method, err := transaction.UnpackInput(inputStr, abiJson)
	if method != "transfer" {
		return info, errors.New("非transfer数据")
	}
	if err != nil {
		return info, errors.New("解析数据失败：" + err.Error())
	}
	if len(unpackData) == 2 {
		p1, ok := unpackData[0].(common.Address)
		if !ok {
			p1, ok = unpackData[1].(common.Address)
			if !ok {
				return info, errors.New("解析收款地址失败")
			}
		}
		info.To = p1.String()
		p2, ok := unpackData[1].(*big.Int)
		if !ok {
			p2, ok = unpackData[0].(*big.Int)
			if !ok {
				return info, errors.New("解析转账金额失败")
			}
		}
		info.Amount = p2
		receiptInfo, err := client.TransactionReceipt(context.Background(), realHash)
		if err != nil {
			return info, err
		}
		if receiptInfo.Status == 1 {
			txMsg, err := trans.AsMessage(types.NewEIP155Signer(chainId), nil)
			if err != nil {
				return info, err
			}
			info.From = txMsg.From().Hex()
			tokenInfo, _ := token2.NewStandardToken(common.HexToAddress(info.Contract), client)
			decimals, _ := tokenInfo.Decimals(nil)
			info.Decimals = decimals
			amountStr := info.Amount.String()
			amountFloat, _ := strconv.ParseFloat(amountStr, 64)
			amountFloat = amountFloat / math.Pow10(int(decimals))
			info.AmountFloat = amountFloat
			return info, nil
		}
	}
	return info, errors.New("非ERC20代币数据")
}

//归集代币
func UnionToken(client *ethclient.Client, authData custom.AuthData, contract, to string, minAmount float64) (*types.Transaction, error) {
	var (
		err error
		tx  *types.Transaction
	)
	if strings.EqualFold(authData.Account, to) {
		return nil, errors.New("付款地址和收款地址一致")
	}
	balance, _ := BalanceOf(client, contract, authData.Account)
	if minAmount > balance {
		return nil, errors.New("余额未达到最少归集量")
	}
	if !account.IsErc20Address(to) {
		return nil, errors.New("非法收款地址")
	}
	token, _ := token2.NewStandardToken(common.HexToAddress(contract), client)
	decimal, _ := token.Decimals(nil)
	balance = balance * math.Pow(10, float64(decimal))
	numStr := strconv.FormatFloat(balance, 'f', 0, 64)
	transAmount, _ := new(big.Int).SetString(numStr, 10)
	toHex := common.HexToAddress(to)
	authOpts, err := account.AuthOpts(client, authData)
	if err != nil {
		return nil, err
	}
	tx, err = token.Transfer(authOpts, toHex, transAmount)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func ToBigAmount(client *ethclient.Client, contract string, amount float64) *big.Int {
	token, _ := token2.NewStandardToken(common.HexToAddress(contract), client)
	decimal, _ := token.Decimals(nil)
	amountDC := dc.NewFromFloat(amount)
	ratioDC := dc.NewFromFloat(math.Pow(10, float64(decimal)))
	bigAmount := amountDC.Mul(ratioDC).BigInt()
	return bigAmount
}
