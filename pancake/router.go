package pancake

import (
	"context"
	"errors"
	"github.com/blueambition/node-lib/account"
	"github.com/blueambition/node-lib/contract/swap/pancake/router"
	token2 "github.com/blueambition/node-lib/contract/token"
	"github.com/blueambition/node-lib/custom"
	"github.com/blueambition/node-lib/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	dc "github.com/shopspring/decimal"
	"math"
	"math/big"
	"time"
)

//Swap Token
func SwapToken(client *ethclient.Client, authData custom.AuthData, approve bool, factoryAddr, routerAddr, t0, t1, receiver string, amountIn float64, minOut float64) (*types.Transaction, error) {
	routerObj, _ := router.NewRouter(common.HexToAddress(routerAddr), client)
	authOpts, err := account.AuthOpts(client, authData)
	if err != nil {
		return nil, err
	}
	pair := GetPair(client, factoryAddr, t0, t1)
	if pair == "" {
		return nil, errors.New("no token pair info")
	}
	if approve {
		//授权gas limit
		authData.GasLimit = 50000
		_, err = token.Approve(client, authData, t0, routerAddr, amountIn)
		if err != nil {
			return nil, err
		}
	}
	//重设nonce
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(authData.Account))
	if err != nil {
		return nil, err
	}
	authOpts.GasLimit = 300000
	authOpts.Nonce = new(big.Int).SetUint64(nonce)
	t0Amount := token.ToBigAmount(client, t0, amountIn)
	t1Amount := token.ToBigAmount(client, t1, minOut)
	path := []common.Address{common.HexToAddress(t0), common.HexToAddress(t1)}
	ts := time.Now().Unix() + 30
	deadline := new(big.Int).SetInt64(ts)
	tx, err := routerObj.SwapExactTokensForTokensSupportingFeeOnTransferTokens(authOpts, t0Amount, t1Amount, path, common.HexToAddress(receiver), deadline)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

//pancake价格
func GetAmountOut(client *ethclient.Client, routerAddr, anchorToken, outToken string) float64 {
	routerObj, _ := router.NewRouter(common.HexToAddress(routerAddr), client)
	tokenObj, _ := token2.NewStandardToken(common.HexToAddress(anchorToken), client)
	decimals, _ := tokenObj.Decimals(nil)
	unit := math.Pow10(int(decimals))
	numBig := dc.NewFromFloat(unit).BigInt()
	path := make([]common.Address, 0)
	path = append(path, common.HexToAddress(anchorToken), common.HexToAddress(outToken))
	price, _ := routerObj.GetAmountsOut(nil, numBig, path)
	if len(price) == 2 {
		outObj, _ := token2.NewStandardToken(common.HexToAddress(outToken), client)
		outDecimals, _ := outObj.Decimals(nil)
		outUnit := dc.NewFromFloat(math.Pow10(int(outDecimals)))
		dcOut, _ := dc.NewFromString(price[1].String())
		outAmount, _ := dcOut.Div(outUnit).RoundFloor(6).Float64()
		return outAmount
	}
	return 0
}
