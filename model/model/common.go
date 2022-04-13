package model

type TxPreExtra struct {
	GasPrice uint64 `json:"gas_price"`
	GasLimit uint64 `json:"gas_limit"`
	Nonce    uint64 `json:"nonce"`
}

// Contract PreCreate Response
type ContractCodeResponse struct {
	TxPreExtra
	Code string `json:"code"`
}
