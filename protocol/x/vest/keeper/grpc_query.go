package keeper

import (
	"context"

	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/x/vest/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VestEntry(
	goCtx context.Context,
	req *types.QueryVestEntryRequest,
) (*types.QueryVestEntryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	// vest 계정에 해당하는 vest entry 가져오기
	vestEntry, err := k.GetVestEntry(ctx, req.VesterAccount)
	if err != nil {
		return nil, err
	}

	// 쿼리에 대한 vest entry 리턴
	return &types.QueryVestEntryResponse{Entry: vestEntry}, nil
}
