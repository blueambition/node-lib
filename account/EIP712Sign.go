package account

import (
	"encoding/json"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

type EIP712Domain struct {
	ChainId           string `json:"chainId"`
	Name              string `json:"name"`
	Version           string `json:"version"`
	VerifyingContract string `json:"verifyingContract"`
}

type EIP712Property struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type EIP712 struct {
	Domain      EIP712Domain `json:"domain"`
	Message     interface{}  `json:"message"`
	PrimaryType string       `json:"primaryType"` //MessagePrimary
	Types       struct {
		EIP712Domain []EIP712Property `json:"EIP712Domain"`
		Primary      []EIP712Property `json:"Primary"`
	} `json:"types"`
}

func ValidEIP712Demo() {
	originData := EIP712{
		Domain: EIP712Domain{"56", "Butterfly", "v1.0", ""},
		Message: struct {
			Content string `json:"content"`
		}{"Bind User"},
		PrimaryType: "Primary",
	}
	originData.Types.EIP712Domain = []EIP712Property{
		{"name", "string"},
		{"version", "string"},
		{"chainId", "uint256"},
		{"verifyingContract", "address"},
	}
	originData.Types.Primary = []EIP712Property{
		{"content", "string"},
	}
	jsonData, _ := json.Marshal(originData)
	fmt.Println(jsonData)
	jsonStr := `{"types":{"EIP712Domain":[{"name":"name","type":"string"},{"name":"version","type":"string"},{"name":"chainId","type":"uint256"}],"Primary":[{"name":"content","type":"string"}]},"domain":{"chainId":"56","name":"Butterfly","version":"v1.0"},"primaryType":"Primary","message":{"content":"Bind User"}}`
	fmt.Println(jsonStr)
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(jsonStr), &typedData); err != nil {
		panic(err)
	}
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress("0x3D12Bd39bB936a73575ea97dFbf308b08b84e76B"),
		typedData,
		"0x5b1a3e3d3f01712ea430333bed3efb70d9f7f6719340bda7f4434260c522f7d72c019ccc40116a3b3c341bb9c5f110683d83874419c101ee1e66c88fe06b4c5c1c",
	)
	fmt.Println(valid, err) // true <nil>
}

func EIP712SignVerify(originData EIP712, sigAddr, signData string) (bool, error) {
	jsonData, _ := json.Marshal(originData)
	var typedData apitypes.TypedData
	if err := json.Unmarshal(jsonData, &typedData); err != nil {
		return false, err
	}
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress(sigAddr),
		typedData,
		signData,
	)
	return valid, err
}
