package backend

import (
	"tokenagent/api"
	"tokenagent/core/initialize"
	"tokenagent/global"
	"tokenagent/rpcserver"

	"go.uber.org/zap"
)

type Backend struct {
	rpcServer *rpcserver.RPCServer
}

func NewBackend() (*Backend, error) {
	backend := &Backend{}
	Router := initialize.Routers()
	backend.rpcServer = rpcserver.NewRPCServer(Router)

	return backend, nil
}

func (backend *Backend) RegisterBackend() error {
	rpc := &api.ApiGroupApp.ChainApi
	collection := &api.ApiGroupApp.CollectionApi
	nft := &api.ApiGroupApp.NftApi
	token := &api.ApiGroupApp.TokenApi

	if err := backend.rpcServer.RegisterName("Collection", collection); err != nil {
		global.MARKET_LOG.Error("", zap.Error(err))
		return err
	}
	if err := backend.rpcServer.RegisterName("Chain", rpc); err != nil {
		global.MARKET_LOG.Error("", zap.Error(err))
		return err
	}
	if err := backend.rpcServer.RegisterName("NFT", nft); err != nil {
		global.MARKET_LOG.Error("", zap.Error(err))
		return err
	}
	if err := backend.rpcServer.RegisterName("Token", token); err != nil {
		global.MARKET_LOG.Error("", zap.Error(err))
		return err
	}

	return nil
}

func (backend *Backend) Start() error {
	if err := backend.rpcServer.Start(); err != nil {
		return err
	}
	return nil
}
