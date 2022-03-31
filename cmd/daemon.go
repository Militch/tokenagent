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
	"os"
	"os/signal"
	"syscall"
	"tokenagent/core"
	"tokenagent/core/backend"
	"tokenagent/global"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:                   "daemon [options]",
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
	Short:                 "Start a tokenagent daemon process",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDaemon()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func runDaemon() error {
	// 初始化Viper
	global.MARKET_VP = core.Viper(cfgFile)
	// 初始化zap日志库
	global.MARKET_LOG = core.Zap()

	var (
		err  error            = nil
		back *backend.Backend = nil
	)

	if back, err = backend.NewBackend(); err != nil {
		return err
	}

	if err = back.RegisterBackend(); err != nil {
		return err
	}

	if err = back.Start(); err != nil {
		return err
	}

	signaleChan := make(chan os.Signal)
	signal.Notify(signaleChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
out:
	select {
	case s := <-signaleChan:
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			break out
		}
	}

	return nil
}
