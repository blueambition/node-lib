package pancake

import (
	"github.com/blueambition/node-lib/contract/swap/pancake/factory"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//获取交易对
func GetPair(client *ethclient.Client, factoryAddr, t0, t1 string) string {
	factoryObj, _ := factory.NewFactory(common.HexToAddress(factoryAddr), client)
	lp, err := factoryObj.GetPair(nil, common.HexToAddress(t0), common.HexToAddress(t1))
	if err != nil {
		return ""
	}
	return lp.Hex()
}
