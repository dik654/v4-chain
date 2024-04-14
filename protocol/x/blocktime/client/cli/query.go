package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryDowntimeParams())
	cmd.AddCommand(CmdQueryAllDowntimeInfo())
	cmd.AddCommand(CmdQueryPreviousBlockInfo())

	return cmd
}

<<<<<<< HEAD
// 아래의 커맨드들은 keeper.go 참조
=======
// down time 관련 설정 파라미터 가져오기
// 세 쿼리 모두 types/query.pb.go에 작성되어있음
>>>>>>> 67021a8090ff8cc0fe64194eb2ac98b2c04335d3
func CmdQueryDowntimeParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-downtime-params",
		Short: "get the DowntimeParams",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.DowntimeParams(
				context.Background(),
				&types.QueryDowntimeParamsRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// 저장되어있는 모든 down time 정보 가져오기
func CmdQueryAllDowntimeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-all-downtime-info",
		Short: "get all downtime info",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.AllDowntimeInfo(
				context.Background(),
				&types.QueryAllDowntimeInfoRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// 이전 블록 정보 가져오기
func CmdQueryPreviousBlockInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-previous-block-info",
		Short: "get previous block info",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.PreviousBlockInfo(
				context.Background(),
				&types.QueryPreviousBlockInfoRequest{},
			)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
