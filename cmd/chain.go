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

// chainCmd represents the chain command
var (
	chainCmd = &cobra.Command{
		Use:                   "chain <command> [options]",
		DisableFlagsInUseLine: true,
		Short:                 "show chain info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	chainSendRawTransactionCmd = &cobra.Command{
		Use:                   "sendrawtransaction [options]",
		DisableFlagsInUseLine: true,
		Short:                 "Send a signed transaction",
		RunE:                  sendRawTransaction,
	}

	chainGetTransactionReceiptCmd = &cobra.Command{
		Use:                   "gettransactionreceipt [options]",
		DisableFlagsInUseLine: true,
		Short:                 "Get transaction receipt",
		RunE:                  getTransactionReceipt,
	}
)

func init() {
	// rootCmd.AddCommand(chainCmd)
	chainCmd.AddCommand(chainSendRawTransactionCmd)
	chainCmd.AddCommand(chainGetTransactionReceiptCmd)
}

func sendRawTransaction(cmd *cobra.Command, args []string) error {
	return nil
}

func getTransactionReceipt(cmd *cobra.Command, args []string) error {
	return nil
}
