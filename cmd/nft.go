/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// nftCmd represents the nft command
var (
	nftCmd = &cobra.Command{
		Use:                   "nft <command> [options]",
		DisableFlagsInUseLine: true,
		Short:                 "show nft info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	nftPreMintCreateCmd = &cobra.Command{
		Use:                   "premintcreate [options]",
		DisableFlagsInUseLine: true,
		Short:                 "Generate transaction body for make NFT",
		RunE:                  preMintCreate,
	}

	nftPreBurnCreateCmd = &cobra.Command{
		Use:                   "preburncreate [options]",
		DisableFlagsInUseLine: true,
		Short:                 "Generate transaction body for destory NFT",
		RunE:                  preBurnCreate,
	}

	nftPreTransferFromCreateCmd = &cobra.Command{
		Use:                   "pretransferfrom [options]",
		DisableFlagsInUseLine: true,
		Short:                 "Generate transaction body for transfer NFT",
		RunE:                  preTransferFromCreate,
	}

	nftOwnerOfCmd = &cobra.Command{
		Use:                   "owner [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get the owner address of NFT",
		RunE:                  ownerOf,
	}

	nftBalanceOfCmd = &cobra.Command{
		Use:                   "balanceof [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get nft balance of address",
		RunE:                  balanceOf,
	}

	nftTokenUriCmd = &cobra.Command{
		Use:                   "tokenuri [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get uri of the specified NFT",
		RunE:                  tokenUri,
	}
)

func init() {
	// rootCmd.AddCommand(nftCmd)
	nftCmd.AddCommand(nftPreMintCreateCmd)
	nftCmd.AddCommand(nftPreBurnCreateCmd)
	nftCmd.AddCommand(nftPreTransferFromCreateCmd)
	nftCmd.AddCommand(nftOwnerOfCmd)
	nftCmd.AddCommand(nftBalanceOfCmd)
	nftCmd.AddCommand(nftTokenUriCmd)

}

func preMintCreate(cmd *cobra.Command, args []string) error {
	return nil
}

func preBurnCreate(cmd *cobra.Command, args []string) error {
	return nil
}

func preTransferFromCreate(cmd *cobra.Command, args []string) error {
	return nil
}

func ownerOf(cmd *cobra.Command, args []string) error {
	return nil
}

func balanceOf(cmd *cobra.Command, args []string) error {
	return nil
}

func tokenUri(cmd *cobra.Command, args []string) error {
	return nil
}
