package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/types"
)

// GetParams returns the DowntimeParams in state.
func (k Keeper) GetDowntimeParams(
	ctx sdk.Context,
) (
	params types.DowntimeParams,
) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(types.DowntimeParamsKey))
	k.cdc.MustUnmarshal(b, &params)
	return params
}

// SetParams updates the Params in state.
// Returns an error iff validation fails.
func (k Keeper) SetDowntimeParams(
	ctx sdk.Context,
	params types.DowntimeParams,
) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set([]byte(types.DowntimeParamsKey), b)

	// For each new duration, we assume the worst case. For new durations that are smaller than all existing
	// durations, we'll use the current block's info. Note that at genesis, this is true for all durations.
	// 새로운 duration이 더 짧으면 기존의 duration이 끝나기 전에 종료되었다는 의미이므로
	// 현재 블록 높이와 Timestamp를 담아서 다시 downtime 시작하기 위한 설정
	newAllDowntimeInfo := types.AllDowntimeInfo{}
	for _, duration := range params.Durations {
		newAllDowntimeInfo.Infos = append(newAllDowntimeInfo.Infos, &types.AllDowntimeInfo_DowntimeInfo{
			Duration: duration,
			BlockInfo: types.BlockInfo{
				Height:    uint32(ctx.BlockHeight()),
				Timestamp: ctx.BlockTime(),
			},
		})
	}

	// Assuming the worst case means assuming that each previously recorded downtime lasted as long as possible.
	// So for each new duration, we take the downtime of the largest existing duration that is smaller.
	// 새로운 duration과 기존의 duration을 비교하여, 새로운 기간이 기존 기간보다 클 때 기존의 정보를 가져오기
	// 즉 기존 duration을 연장한다고 생각
	allDowntimeInfo := k.GetAllDowntimeInfo(ctx)
	for _, info := range newAllDowntimeInfo.Infos {
		for _, oldInfo := range allDowntimeInfo.Infos {
			if info.Duration >= oldInfo.Duration {
				info.BlockInfo = oldInfo.BlockInfo
			} else {
				break
			}
		}
	}
	// keeper KVStore에 저장
	k.SetAllDowntimeInfo(ctx, &newAllDowntimeInfo)
	return nil
}
