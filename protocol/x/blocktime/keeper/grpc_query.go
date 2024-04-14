package keeper

import (
	"context"

	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/types"

	// 에러코드
	"google.golang.org/grpc/codes"
	// 리턴용 status
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// keeper.go의 메서드 gRPC 쿼리에 대한 respose들
func (k Keeper) DowntimeParams(
	c context.Context,
	req *types.QueryDowntimeParamsRequest,
) (
	*types.QueryDowntimeParamsResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	params := k.GetDowntimeParams(ctx)
	return &types.QueryDowntimeParamsResponse{
		Params: params,
	}, nil
}

func (k Keeper) PreviousBlockInfo(
	c context.Context,
	req *types.QueryPreviousBlockInfoRequest,
) (
	*types.QueryPreviousBlockInfoResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	info := k.GetPreviousBlockInfo(ctx)
	return &types.QueryPreviousBlockInfoResponse{
		Info: &info,
	}, nil
}

func (k Keeper) AllDowntimeInfo(
	c context.Context,
	req *types.QueryAllDowntimeInfoRequest,
) (
	*types.QueryAllDowntimeInfoResponse,
	error,
) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := lib.UnwrapSDKContext(c, types.ModuleName)
	info := k.GetAllDowntimeInfo(ctx)
	return &types.QueryAllDowntimeInfoResponse{
		Info: info,
	}, nil
}
