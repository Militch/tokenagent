package tokenapi

import (
	"math/big"
	"tokenagent/model/model"
	"tokenagent/service"
)

type TokenAPIHandler struct {
}

//@author: [libofeng]
//@function: PreCreate
//@description: 生成代币的合约码，返回给前端进行签名
//@param: args model.TokenPreCreateRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreCreate(args model.TokenPreCreateRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: Name
//@description: 获取代币的名称，直接调用无需签名
//@param: args model.TokenRequest
//@return: resp *string
func (handler *TokenAPIHandler) Name(args model.TokenCallRequest, resp *string) error {
	response, err := service.ServiceGroupApp.TokenProxy.Name(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: Symbol
//@description: 获取代币的符号，直接调用无需签名
//@param: args model.CollectionRequest
//@return: resp *string
func (handler *TokenAPIHandler) Symbol(args model.TokenCallRequest, resp *string) error {
	response, err := service.ServiceGroupApp.TokenProxy.Symbol(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: TotalSupply
//@description: 获取发行代币的总额，直接调用无需签名
//@param: args model.TokenRequest
//@return:  resp **big.Int
func (handler *TokenAPIHandler) TotalSupply(args model.TokenCallRequest, resp **big.Int) error {
	response, err := service.ServiceGroupApp.TokenProxy.TotalSupply(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: BalanceOf
//@description: 获取发行代币的总额，直接调用无需签名
//@param: args model.TokenBalanceOfCallRequest
//@return:  resp **big.Int
func (handler *TokenAPIHandler) BalanceOf(args model.TokenBalanceOfCallRequest, resp **big.Int) error {
	response, err := service.ServiceGroupApp.TokenProxy.BalanceOf(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreMint
//@description: 生成发行代币的code码，返回给前端进行签名
//@param: args model.TokenPreMintRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreMint(args model.TokenPreMintRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreMintCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreApprove
//@description: 生成授权代币的code码，返回给前端进行签名
//@param: args model.TokenPreApproveRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreApprove(args model.TokenPreApproveRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreApproveCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreBurnFrom
//@description: 生成销毁代币的code码，返回给前端进行签名
//@param: args model.TokenPreBurnRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreBurnFrom(args model.TokenPreBurnRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreBurnFromCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreIncreaseAllowance
//@description: 生成添加津贴的code码，返回给前端进行签名
//@param: args model.TokenPreAllowanceRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreIncreaseAllowance(args model.TokenPreAllowanceRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreIncreaseAllowance(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreDecreaseAllowance
//@description: 生成减少津贴的code码，返回给前端进行签名
//@param: args model.TokenPreAllowanceRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreDecreaseAllowance(args model.TokenPreAllowanceRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreDecreaseAllowance(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreTransferFrom
//@description: 生成销毁代币的code码，返回给前端进行签名
//@param: args model.TokenPreTransferRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreTransferFrom(args model.TokenPreTransferRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreTransferFromCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PrePause
//@description: 生成暂停代币的code码，返回给前端进行签名
//@param: args model.TokenPreRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PrePause(args model.TokenPreRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPrePauseCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}

//@author: [libofeng]
//@function: PreUnPause
//@description: 生成取消暂停代币的code码，返回给前端进行签名
//@param: args model.TokenPreRequest
//@return: resp **model.ContractCodeResponse
func (handler *TokenAPIHandler) PreUnPause(args model.TokenPreRequest, resp **model.ContractCodeResponse) error {
	response, err := service.ServiceGroupApp.TokenProxy.TokenPreUnPauseCreate(args)
	if err != nil {
		return err
	}
	*resp = response
	return nil
}
