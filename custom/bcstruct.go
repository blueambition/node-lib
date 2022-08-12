package custom

import "math/big"

//认证信息传值
type AuthData struct {
	KsPath      string
	Account     string
	Password    string
	GasLimit    uint64
	GasMultiple float64
	ETHValue    *big.Int
}

//转账信息传值
type TransData struct {
	KsPath   string  `json:"ks_path"`
	From     string  `json:"from"`
	Password string  `json:"password"`
	To       string  `json:"to"`
	Amount   float64 `json:"amount"`
	GasLimit uint64  `json:"gas_limit"`
}

//归集信息传值
type UnionData struct {
	KsPath    string  `json:"ks_path"`
	From      string  `json:"from"`
	Password  string  `json:"password"`
	To        string  `json:"to"`
	MinAmount float64 `json:"min_amount"`
	FeeAmount float64 `json:"fee_amount"`
	GasLimit  uint64  `json:"gas_limit"`
}

//代币交易信息解析
type TransTokenInfo struct {
	From        string   `json:"from"`
	Contract    string   `json:"contract"`
	To          string   `json:"to"`
	Amount      *big.Int `json:"amount"`
	Decimals    uint8    `json:"decimal"`
	AmountFloat float64  `json:"amount_float"`
}

//跨链桥交易信息解析
type BridgeTransInfo struct {
	From        string   `json:"from"`
	Contract    string   `json:"contract"`
	Token       string   `json:"token"`
	Amount      *big.Int `json:"amount"`
	Decimals    uint8    `json:"decimal"`
	AmountFloat float64  `json:"amount_float"`
}
