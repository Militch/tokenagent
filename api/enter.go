package api

import (
	"tokenagent/api/chainapi"
	"tokenagent/api/collectionapi"
	"tokenagent/api/nftapi"
	"tokenagent/api/tokenapi"
)

type ApiGroup struct {
	ChainApi      chainapi.ChainAPIHandler
	CollectionApi collectionapi.CollectionAPIHandler
	NftApi        nftapi.NftAPIHandler
	TokenApi      tokenapi.TokenAPIHandler
}

var ApiGroupApp = new(ApiGroup)
