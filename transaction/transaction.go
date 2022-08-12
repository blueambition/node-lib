package transaction

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/blueambition/node-lib/account"
	"github.com/blueambition/node-lib/custom"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
	"strconv"
	"strings"
)

//16进制转10进制
func HexToInt(hex string) int {
	hex = strings.Replace(hex, "0x", "", 1)
	n, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}

// BigInt转成16进制
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}
	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}

// int转成16进制
func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}

//解析输入
func UnpackInput(txInput, abiJson string) ([]interface{}, string, error) {
	var data = make([]interface{}, 0)
	abi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return data, "", err
	}
	if len(txInput) > 10 {
		// decode txInput method signature
		decodeSign, err := hex.DecodeString(txInput[2:10])
		if err != nil {
			return data, "", err
		}
		// recover Method from signature and ABI
		method, err := abi.MethodById(decodeSign)
		if err != nil {
			return data, "", err
		}
		//decodeData, err := hex.DecodeString(txInput[2:])
		decodedData, err := hex.DecodeString(txInput[10:])
		if err != nil {
			return data, "", err
		}
		// unpack method inputs
		data, err = method.Inputs.Unpack(decodedData)
		return data, method.Name, err
	}
	return data, "", errors.New("数据：" + txInput + "解析失败")
}

//转账
func SendTransaction(client *ethclient.Client, transData custom.TransData) (*types.Transaction, error) {
	if !account.IsErc20Address(transData.To) {
		return nil, errors.New("非法收款地址")
	}
	if strings.EqualFold(transData.From, transData.To) {
		return nil, errors.New("付款地址和收款地址一致")
	}
	balance := account.Balance(client, transData.From)
	if transData.Amount > balance {
		return nil, errors.New("余额不足")
	}
	priKeyStr, err := account.GetPriKey(transData.KsPath, transData.From, transData.Password)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(priKeyStr)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("公钥解密失败")
	}
	fromHex := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromHex)
	if err != nil {
		return nil, err
	}
	transData.Amount = transData.Amount * math.Pow(10, 18)
	transAmountStr := strconv.FormatFloat(transData.Amount, 'f', 0, 64)
	value, _ := new(big.Int).SetString(transAmountStr, 10)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	toHex := common.HexToAddress(transData.To)
	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      transData.GasLimit,
		Value:    value,
		To:       &toHex,
	}
	tx := types.NewTx(&txData)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return nil, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	//txHash := signedTx.Hash().Hex()
	return signedTx, nil
}

//归集
func UnionTransaction(client *ethclient.Client, unionData custom.UnionData) (*types.Transaction, error) {
	if !account.IsErc20Address(unionData.To) {
		return nil, errors.New("非法收款地址")
	}
	if strings.EqualFold(unionData.From, unionData.To) {
		return nil, errors.New("付款地址和收款地址一致")
	}
	balance := account.Balance(client, unionData.From)
	if unionData.MinAmount > balance {
		return nil, errors.New("余额未达到最小归集量")
	}
	amount := balance - unionData.FeeAmount
	priKeyStr, err := account.GetPriKey(unionData.KsPath, unionData.From, unionData.Password)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(priKeyStr)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("公钥解密失败")
	}
	fromHex := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromHex)
	if err != nil {
		return nil, err
	}
	amount = amount * math.Pow(10, 18)
	amountStr := strconv.FormatFloat(amount, 'f', 0, 64)
	value, _ := new(big.Int).SetString(amountStr, 10)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	toHex := common.HexToAddress(unionData.To)
	txData := types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      unionData.GasLimit,
		Value:    value,
		To:       &toHex,
	}
	tx := types.NewTx(&txData)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		return nil, err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}
	//txHash := signedTx.Hash().Hex()
	return signedTx, nil
}

//区块交易信息
func GetTransByBlockNo(client *ethclient.Client, blockNo uint64) types.Transactions {
	blocks, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNo))
	if err != nil {
		return nil
	}
	return blocks.Transactions()
}

//获取交易回执单
func GetTransReceipt(client *ethclient.Client, txHash string) (*types.Receipt, error) {
	realHash := common.HexToHash(txHash)
	receiptInfo, err := client.TransactionReceipt(context.Background(), realHash)
	return receiptInfo, err
}

//验证Hash
func IsTxHash(client *ethclient.Client, txHash string) error {
	realHash := common.HexToHash(txHash)
	_, isPending, err := client.TransactionByHash(context.Background(), realHash)
	if err != nil {
		return err
	}
	if isPending {
		return errors.New("tx is pending status")
	}
	return nil
}
