package cmd

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/dydxprotocol/v4-chain/protocol/daemons/configs"
	"github.com/spf13/cobra"
)

// AddInitCmdPostRunE adds a PostRunE to the `init` subcommand.
func AddInitCmdPostRunE(rootCmd *cobra.Command) {
	// Fetch init subcommand.
	initCmd, _, err := rootCmd.Find([]string{"init"})
	if err != nil {
		os.Exit(1)
	}

	// Add PostRun to configure required setups after `init`.
	initCmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		// Get home directory.
		// cmd에서 client.context에 해당하는 컨텍스트가 존재한다면 해당 컨텍스트를 리턴
		// 없다면 새로운 컨텍스트를 생성하여 리턴
		clientCtx := client.GetClientContextFromCmd(cmd)

		// Add default pricefeed exchange config toml file if it does not exist.
		// pricefeed exchange toml이 없다면 홈 디렉터리를 포함하여 생성
		configs.WriteDefaultPricefeedExchangeToml(clientCtx.HomeDir)
		return nil
	}
}
