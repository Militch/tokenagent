package model

type TxPreExtra struct {
	GasPrice uint64 `json:"gas_price"`
	GasLimit uint64 `json:"gas_limit"`
	Nonce    uint64 `json:"nonce"`
}
