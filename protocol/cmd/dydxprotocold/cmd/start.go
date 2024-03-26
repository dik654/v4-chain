package cmd

import (
	appflags "github.com/dydxprotocol/v4-chain/protocol/app/flags"
	daemonflags "github.com/dydxprotocol/v4-chain/protocol/daemons/flags"
	"github.com/dydxprotocol/v4-chain/protocol/indexer"
	clobflags "github.com/dydxprotocol/v4-chain/protocol/x/clob/flags"
	"github.com/spf13/cobra"
)

// GetOptionWithCustomStartCmd returns a root command option with custom start commands.
func GetOptionWithCustomStartCmd() *RootCmdOption {
	option := newRootCmdOption()
	f := func(cmd *cobra.Command) {
		// Add app flags.
		// app/flags/flags.go
		// 호스트명, 포트, GRPC 여부, 풀노드 여부 관련
		// full node start 명령어의 flag argument 선언
		appflags.AddFlagsToCmd(cmd)

		// Add daemon flags.
		// daemons/flags/flags.go
		// price, bridge, liquidation oracle 관련
		// full node start 명령어의 flag argument 선언
		daemonflags.AddDaemonFlagsToCmd(cmd)

		// Add indexer flags.
		// indexer/flags.go
		// full node start 명령어의 카프카 관련 flag argument 선언
		indexer.AddIndexerFlagsToCmd(cmd)

		// Add clob flags.
		// x/clob/flags/flags.go
		// full node start 명령어의 청산시도 횟수, Telemetry 관련 flag argument 선언
		clobflags.AddClobFlagsToCmd(cmd)
	}
	// cmd/dydxprotocold/cmd/root_option.go
	// 위의 flag argument들을 start에 등록
	option.setCustomizeStartCmd(f)
	return option
}
