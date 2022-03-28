package nftservice

import (
	"tokenagent/model/model"
)

type XFSNft struct {
}

func (xfs *XFSNft) NftPreMintCreate(args model.NftPreMintRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreApproveCreate(args model.NftPreApproveRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreBurnCreate(args model.NftPreBurnRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreSetApprovalForAllCreate(args model.NftPreSetApprovalForAllRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreTransferOwnershipCreate(args model.NftPreTransferOwnershipRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) OwnerOf(args model.NftOwnerOfRequest) (string, error) {
	return "", nil
}

func (xfs *XFSNft) BalanceOf(args model.NftBalanceOfRequest) (uint64, error) {
	return uint64(0), nil
}

func (xfs *XFSNft) TokenUri(args model.NftTokenInfoRequest) (string, error) {
	return "", nil
}

func (xfs *XFSNft) NftPreSafeTransferFromCreate(args model.NftPreTransferFromRequest) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) NftPreSafeTransferFromExCreate(args model.NftPreTransferFromRequestEx) (*model.ContractCodeResponse, error) {
	return nil, nil
}

func (xfs *XFSNft) GetApproved(args model.NftTokenInfoRequest) (string, error) {
	return "", nil
}
