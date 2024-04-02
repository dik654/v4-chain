package blocktime

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/keeper"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/types"
)

// InitGenesis initializes the blocktime module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

 // genState.Params는 types/genesis.pb.go에 선언되어있음
// gRPC용
	if err := k.SetDowntimeParams(ctx, genState.Params); err != nil {
		panic(err)
	}

// Set to genesis block height and time
// context에 담긴 현재시간과 블록높이로 블록정보 구조체 생성
	k.SetPreviousBlockInfo(ctx, &types.BlockInfo{
		Height:    uint32(ctx.BlockHeight()),
		Timestamp: ctx.BlockTime(),
	})
}

// ExportGenesis returns the blocktime module's exported genesis.
// InitGenesis에서 설정한 GenesisState 리턴
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetDowntimeParams(ctx),
	}
}
