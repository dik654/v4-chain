package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListSubaccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-subaccount",
		Short: "list all subaccount",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			// pagination page 요청 구조체 받기
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSubaccountRequest{
				Pagination: pageReq,
			}

			// grpc_query_subaccount.go
			// 모든 sub account 가져오기
			res, err := queryClient.SubaccountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSubaccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-subaccount [index]",
		Short: "shows a subaccount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			// gRPC 쿼리용 client 객체 생성
			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]
			// 2번째 인수를 uint32로 변환
			argNumber, err := cast.ToUint32E(args[1])
			if err != nil {
				return err
			}

			params := &types.QueryGetSubaccountRequest{
				Owner:  argOwner,
				Number: argNumber,
			}

			// 해당 유저의 sub account 가져오기
			res, err := queryClient.Subaccount(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
