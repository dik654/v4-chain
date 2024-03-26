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
		appflags.AddFlagsToCmd(cmd)

		// Add daemon flags.
		// daemons/flags/flags.go
		daemonflags.AddDaemonFlagsToCmd(cmd)

		// Add indexer flags.
		// indexer/flags.go
		indexer.AddIndexerFlagsToCmd(cmd)

		// Add clob flags.
		// x/clob/flags/flags.go
		clobflags.AddClobFlagsToCmd(cmd)
	}
	// cmd/dydxprotocold/cmd/root_option.go
	option.setCustomizeStartCmd(f)
	return option
}
