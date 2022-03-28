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

// collectionCmd represents the collection command
var (
	collectionCmd = &cobra.Command{
		Use:                   "collection <command> [options]",
		DisableFlagsInUseLine: true,
		Short:                 "show collection info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	collectionPreCreateCmd = &cobra.Command{
		Use:                   "precreate [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get collection tx body for sig",
		RunE:                  preCreate,
	}

	collectionNameCmd = &cobra.Command{
		Use:                   "name [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get collection name",
		RunE:                  name,
	}

	collectionSymbolCmd = &cobra.Command{
		Use:                   "symbol [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get collection symbol",
		RunE:                  symbol,
	}

	collectionContractURICmd = &cobra.Command{
		Use:                   "contracturi [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get collection contracturi",
		RunE:                  contractURI,
	}

	collectionTokenURIPrefixCmd = &cobra.Command{
		Use:                   "collectionPreCreateCmd [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get collection tokenuri prefix",
		RunE:                  tokenURIPrefix,
	}

	collectionOwnerCmd = &cobra.Command{
		Use:                   "owner [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get the owner address of collection",
		RunE:                  owner,
	}

	collectionTotalSupplyCmd = &cobra.Command{
		Use:                   "totalsupply [options]",
		DisableFlagsInUseLine: true,
		Short:                 "get the total supply of collection",
		RunE:                  totalSupply,
	}
)

func init() {
	// rootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionPreCreateCmd)
	collectionCmd.AddCommand(collectionNameCmd)
	collectionCmd.AddCommand(collectionSymbolCmd)
	collectionCmd.AddCommand(collectionContractURICmd)
	collectionCmd.AddCommand(collectionTokenURIPrefixCmd)
	collectionCmd.AddCommand(collectionOwnerCmd)
	collectionCmd.AddCommand(collectionTotalSupplyCmd)
}

func preCreate(cmd *cobra.Command, args []string) error {
	return nil
}

func name(cmd *cobra.Command, args []string) error {
	return nil
}

func symbol(cmd *cobra.Command, args []string) error {
	return nil
}

func contractURI(cmd *cobra.Command, args []string) error {
	return nil
}

func tokenURIPrefix(cmd *cobra.Command, args []string) error {
	return nil
}

func owner(cmd *cobra.Command, args []string) error {
	return nil
}

func totalSupply(cmd *cobra.Command, args []string) error {
	return nil
}
