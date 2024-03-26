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
	// start에 flag argument 등록
	option := cmd.GetOptionWithCustomStartCmd()
	// cmd/dydxprotocold/cmd/root.go
	// dydxprotocold를 실행시키면 나오는 명령어 리스트 생성
	rootCmd := cmd.NewRootCmd(option, app.DefaultNodeHome)

	// cmd/dydxprotocold/cmd/tendermint.go
	// 개인키 생성 명령어 및 개인키, 코덱등 디버그 명령어 생성
	cmd.AddTendermintSubcommands(rootCmd)
	// cmd/dydxprotocold/cmd/init.go
	// init 명령어 생성
	cmd.AddInitCmdPostRunE(rootCmd)

	if err := svrcmd.Execute(rootCmd, constants.AppDaemonName, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
