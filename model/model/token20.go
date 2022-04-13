package model

type TokenPreCreateRequest struct {
	BlockChain  string `json:"blockchain"`
	FromAddress string `json:"from_address"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}

// ERC20的基本信息请求
type TokenCallRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
}

type TokenBalanceOfCallRequest struct {
	BlockChain      string `json:"blockchain"`
	ContractAddress string `json:"contract_address"`
	Account         string `json:"account"`
}

type TokenPreMintRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
	ToAddress       string `json:"to_address"`
	Amount          string `json:"amount"`
}

type TokenPreApproveRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
	SpenderAddress  string `json:"spender_address"`
	Amount          string `json:"amount"`
}

type TokenPreBurnRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
	Amount          string `json:"amount"`
}

type TokenPreTransferRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
	ToAddress       string `json:"to_address"`
	Amount          string `json:"amount"`
}

type TokenPreAllowanceRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
	SpenderAddress  string `json:"spender_address"`
	Value           string `json:"value"`
}

type TokenPreRequest struct {
	BlockChain      string `json:"blockchain"`
	FromAddress     string `json:"from_address"`
	ContractAddress string `json:"contract_address"`
}
