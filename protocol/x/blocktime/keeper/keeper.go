package keeper

import (
	"fmt"
	"time"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/lib/metrics"
	"github.com/dydxprotocol/v4-chain/protocol/x/blocktime/types"
)

// 키퍼에는 바이너리 직렬화 코덱, keeper의 KVStore에 접근하기위한 StoreKey, whitelist
type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		authorities map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		authorities: lib.UniqueSliceToSet(authorities),
	}
}

func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {}

// downtime은 마지막 블록시간과 현재 블록시간의 차이
func (k Keeper) GetAllDowntimeInfo(ctx sdk.Context) *types.AllDowntimeInfo {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get([]byte(types.AllDowntimeInfoKey))

	if bytes == nil {
		return &types.AllDowntimeInfo{}
	}

	var info types.AllDowntimeInfo
	k.cdc.MustUnmarshal(bytes, &info)
	return &info
}

// SetAllDowntimeInfo sets AllDowntimeInfo in state. Durations in AllDowntimeInfo must match
// the durations in DowntimeParams. If not, behavior of this module is undefined.
func (k Keeper) SetAllDowntimeInfo(ctx sdk.Context, info *types.AllDowntimeInfo) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(info)
	store.Set([]byte(types.AllDowntimeInfoKey), b)
}

// types/blocktime.pb.go에 types.BlockInfo 구조체가 선언되어있음
// keeper의 KVstore에 저장된 마샬링된 데이터 언마샬링 후 리턴
func (k Keeper) GetPreviousBlockInfo(ctx sdk.Context) types.BlockInfo {
	store := ctx.KVStore(k.storeKey)
	bytes := store.Get([]byte(types.PreviousBlockInfoKey))

	if bytes == nil {
		return types.BlockInfo{}
	}

	var info types.BlockInfo
	k.cdc.MustUnmarshal(bytes, &info)
	return info
}

// GetTimeSinceLastBlock returns the time delta between the current block time and the last block time.
// 언마샬링하여 받은 이전 블록 구조체에서 시간을 가져와서 현재시간과 차이 계산
func (k Keeper) GetTimeSinceLastBlock(ctx sdk.Context) time.Duration {
	prevBlockInfo := k.GetPreviousBlockInfo(ctx)
	return ctx.BlockTime().Sub(prevBlockInfo.Timestamp)
}

// 현재 블록 시간과 높이를 구조체에 담아서 마샬링
// keeper의 KVStore에 저장
func (k Keeper) SetPreviousBlockInfo(ctx sdk.Context, info *types.BlockInfo) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(info)
	store.Set([]byte(types.PreviousBlockInfoKey), b)
}

// UpdateAllDowntimeInfo updates AllDowntimeInfo by considering the downtime between the current block and
// the previous block and updating the DowntimeInfo for each observed duration.
// 이전 블록시간과 현재 블록시간과의 차이(지난시간)보다 duration이 작은 경우(아직 경과시간이 지나지 않았다면)
// 블록높이와 블록시간을 현재의 것으로 변경
func (k Keeper) UpdateAllDowntimeInfo(ctx sdk.Context) {
	previousBlockInfo := k.GetPreviousBlockInfo(ctx)
	// 마지막 블록시간과 현재시간의 차이
	delta := ctx.BlockTime().Sub(previousBlockInfo.Timestamp)
	// Report block time in milliseconds.
	// prometheus로 보내기
	telemetry.SetGauge(
		float32(delta.Milliseconds()),
		types.ModuleName,
		metrics.BlockTimeMs,
	)

	// 모든 downtime 정보를 가져와서
	allInfo := k.GetAllDowntimeInfo(ctx)

	// 업데이트 duration 간격을 넘었을 경우 업데이트
	for _, info := range allInfo.Infos {
		if delta >= info.Duration {
			info.BlockInfo = types.BlockInfo{
				Height:    uint32(ctx.BlockHeight()),
				Timestamp: ctx.BlockTime(),
			}
		} else {
			break
		}
	}
	k.SetAllDowntimeInfo(ctx, allInfo)
}

// GetDowntimeInfoFor gets the DowntimeInfo for a specific duration. If the exact duration is not observed, it
// returns the DowntimeInfo for the largest duration that is smaller than the input duration. If the input
// duration is smaller than all observed durations, then return a DowntimeInfo with duration 0 and the current
// block's info.
// 입력 duration보다 작은 duration중 가장 큰 duration을 가진 downtime info 리턴
// 해당하는게 없다면 duration 0에 현재시간을 갖는 downtime info 리턴
func (k Keeper) GetDowntimeInfoFor(ctx sdk.Context, duration time.Duration) types.AllDowntimeInfo_DowntimeInfo {
	allInfo := k.GetAllDowntimeInfo(ctx)
	ret := types.AllDowntimeInfo_DowntimeInfo{
		Duration: 0,
		BlockInfo: types.BlockInfo{
			Height:    uint32(ctx.BlockHeight()),
			Timestamp: ctx.BlockTime(),
		},
	}
	// Duration이 지난 downtime들만 리턴
	for _, info := range allInfo.Infos {
		if duration >= info.Duration {
			ret = *info
		} else {
			break
		}
	}
	return ret
}
