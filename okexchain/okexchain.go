package okexchain

import (
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/okex/exchain/app"
	"github.com/okex/exchain/app/codec"
	evmtypes "github.com/okex/exchain/x/evm/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/bytes"
)

//获取OKHash
func GetTxHash(signedTx *types.Transaction) string {
	//ts := types.Transactions{signedTx}
	//rawTx := hex.EncodeToString(ts)
	rawTx, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return ""
	}
	tx := new(evmtypes.MsgEthereumTx)
	// RLP decode raw transaction bytes
	if err = rlp.DecodeBytes(rawTx, tx); err != nil {
		return ""
	}
	cdc := codec.MakeCodec(app.ModuleBasics)
	txEncoder := authclient.GetTxEncoder(cdc)
	txBytes, err := txEncoder(tx)
	if err != nil {
		return ""
	}
	var hexBytes bytes.HexBytes
	hexBytes = tmhash.Sum(txBytes)
	hash := common.HexToHash(hexBytes.String())
	return hash.String()
}
