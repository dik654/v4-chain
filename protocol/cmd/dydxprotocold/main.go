package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/dydxprotocol/v4-chain/protocol/app"
	"github.com/dydxprotocol/v4-chain/protocol/app/config"
	"github.com/dydxprotocol/v4-chain/protocol/app/constants"
	"github.com/dydxprotocol/v4-chain/protocol/cmd/dydxprotocold/cmd"
)

func main() {
	// app/config/config.go
	// config 구조체 생성 및 고정
	// prefix 설정
	config.SetupConfig()

	// cmd/dydxprotocold/cmd/start.go
	option := cmd.GetOptionWithCustomStartCmd()
	// cmd/dydxprotocold/cmd/root.go
	rootCmd := cmd.NewRootCmd(option, app.DefaultNodeHome)

	// cmd/dydxprotocold/cmd/tendermint.go
	cmd.AddTendermintSubcommands(rootCmd)
	// cmd/dydxprotocold/cmd/init.go
	cmd.AddInitCmdPostRunE(rootCmd)

	if err := svrcmd.Execute(rootCmd, constants.AppDaemonName, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
