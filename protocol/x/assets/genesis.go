package assets

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/x/assets/keeper"
	"github.com/dydxprotocol/v4-chain/protocol/x/assets/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// 체인 초기화 전 constructor 역할
	k.InitializeForGenesis(ctx)

	// 체인 asset keeper 구조체 생성 
	for _, asset := range genState.Assets {
		_, err := k.CreateAsset(
			ctx,
			asset.Id,
			asset.Symbol,
			asset.Denom,
			asset.DenomExponent,
			asset.HasMarket,
			asset.MarketId,
			asset.AtomicResolution,
		)
		if err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// 생성한 체인의 genesis를 JSON파일로 마샬링
	genesis := types.DefaultGenesis()
	// 마샬링한 JSON파일에 현재 존재하는 asset 데이터들 추가
	genesis.Assets = k.GetAllAssets(ctx)
	return genesis
}
