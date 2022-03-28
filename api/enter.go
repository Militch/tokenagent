package api

import (
	"tokenagent/api/chainapi"
	"tokenagent/api/collectionapi"
	"tokenagent/api/nftapi"
)

type ApiGroup struct {
	ChainApi      chainapi.ChainAPIHandler
	CollectionApi collectionapi.CollectionAPIHandler
	NftApi        nftapi.NftAPIHandler
}

var ApiGroupApp = new(ApiGroup)
