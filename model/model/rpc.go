package model

import (
	"math/big"
)

// RPC Request
type TxRequest struct {
	BlockChain string `json:"blockchain"`
	Data       string `json:"data"`
}

// Tx Response
type TxResponse struct {
	TxHash          string `json:"txhash"`
	ContractAddress string `json:"contract_address"`
}

type TxRecRequest struct {
	BlockChain      string `json:"blockchain"`
	TxHash          string `json:"txhash"`
	ContractAddress string `json:"contract_address"`
}

type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8  `json:"type,omitempty"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          string `json:"transactionHash" gencodec:"required"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         uint64 `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        string   `json:"blockHash,omitempty"`
	BlockNumber      *big.Int `json:"blockNumber,omitempty"`
	TransactionIndex uint     `json:"transactionIndex"`
}
