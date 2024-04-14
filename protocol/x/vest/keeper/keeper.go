package keeper

import (
	"fmt"
	"math/big"
	"time"

	errorsmod "cosmossdk.io/errors"
	cosmoslog "cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/lib/log"
	"github.com/dydxprotocol/v4-chain/protocol/lib/metrics"
	"github.com/dydxprotocol/v4-chain/protocol/x/vest/types"
	gometrics "github.com/hashicorp/go-metrics"
)

type (
	Keeper struct {
		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		bankKeeper      types.BankKeeper
		blockTimeKeeper types.BlockTimeKeeper
		authorities     map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	blockTimeKeeper types.BlockTimeKeeper,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		bankKeeper:      bankKeeper,
		blockTimeKeeper: blockTimeKeeper,
		authorities:     lib.UniqueSliceToSet(authorities),
	}
}

// 등록된 모듈 주소만 vest 실행가능
func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) Logger(ctx sdk.Context) cosmoslog.Logger {
	return ctx.Logger().With(cosmoslog.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

// Process vesting for all vest entries. Intended to be called in BeginBlocker.
// For each vest entry:
// 1. Return if `block_time <= vest_entry.start_time` (vesting has not started yet)
// 2. Return if `prev_block_time >= vest_entry.end_time` (vesting has ended)
// 3. Transfer the following amount of tokens from vester account to treasury account:
//
//		  min(
//			(block_time - last_vest_time) / (end_time - last_vest_time),
//		 	1
//		  ) * vester_account_balance
//
//	  where `last_vest_time = max(start_time, prev_block_time)`
//
// 매 블록 시작마다 처리
func (k Keeper) ProcessVesting(ctx sdk.Context) {
	// Convert timestamps to milliseconds for algebraic operations.
	// 시간을 unix 시간으로 변환
	blockTimeMilli := ctx.BlockTime().UnixMilli()
	// 이전 블록정보와 timestamp 가져오기
	prevBlockInfo := k.blockTimeKeeper.GetPreviousBlockInfo(ctx)
	prevBlockTimeMilli := prevBlockInfo.Timestamp.UnixMilli()

	// Process each vest entry.
	// 모든 entry를 돌면서
	for _, entry := range k.GetAllVestEntries(ctx) {
		// 시작 종료시간 가져와서
		startTimeMilli := entry.StartTime.UnixMilli()
		endTimeMilli := entry.EndTime.UnixMilli()
		// `block_time` <= `start_time`. Vesting has not started.
		// 현재 시간이 시작 시간보다 작다면 pass
		if blockTimeMilli <= startTimeMilli {
			continue
		}
		// 이전 블록시간에 이미 종료 시간이 지난시점이라면 pass
		// `end_time` <= `prev_block_time`. Vesting has ended.
		if endTimeMilli <= prevBlockTimeMilli {
			continue
		}

		// 마지막으로 vesting 시간은 초기값은 vesting 시작시간
		// 시작 이후 접근한 적이 있다면 이전 블록시간(마지막으로 vesting을 실행한 시간)
		lastVestTimeMilli := lib.Max(startTimeMilli, prevBlockTimeMilli)
		// 경과시간 / 남은시간 계산
		bigRatVestProportion := big.NewRat(blockTimeMilli-lastVestTimeMilli, endTimeMilli-lastVestTimeMilli)

		// vester의 주소의 balance 확인
		vesterBalance := k.bankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(entry.VesterAccount), entry.Denom)
		vestAmount := vesterBalance.Amount
		// 경과시간이 남은시간보다 작다면
		if bigRatVestProportion.Cmp(lib.BigRat1()) < 0 {
			// vest_amount = vester_balance * vestProportion
			bigRatBalance := new(big.Rat).SetInt(vesterBalance.Amount.BigInt())
			bigRatVestAmount := new(big.Rat).Mul(
				bigRatBalance,
				bigRatVestProportion,
			)
			// 권리가 해제된 amount
			vestAmount = sdkmath.NewIntFromBigInt(lib.BigRatRound(bigRatVestAmount, false))
		}

		// 해제된 amount가 존재한다면
		if !vestAmount.IsZero() {
			// treasury account로 해제된 amount만큼 전송
			if err := k.bankKeeper.SendCoinsFromModuleToModule(
				ctx,
				entry.VesterAccount,
				entry.TreasuryAccount,
				sdk.NewCoins(sdk.NewCoin(entry.Denom, vestAmount)),
			); err != nil {
				// transfer 과정에서 문제가 있다면 종료시키지않고 에러 log만 생성
				// This should never happen. However, if it does, we should not panic.
				// ProcessVesting is called in BeginBlocker, and panicking in BeginBlocker could cause liveness issues.
				// Instead, we generate an informative error log, emit an error metric, and continue.
				log.ErrorLogWithError(
					ctx,
					"unexpected internal error: failed to transfer vest amount to treasury account",
					err,
					"vester_account",
					entry.VesterAccount,
					"treasury_account",
					entry.TreasuryAccount,
					"denom",
					entry.Denom,
					"vest_amount",
					vestAmount,
					"vest_account_balance",
					vesterBalance,
				)
				// 에러 counter 증가시키기
				telemetry.IncrCounter(1, metrics.ProcessVesting, metrics.AccountTransfer, metrics.Error)
				continue
			}
		}

		// event 발생
		telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, metrics.VestAmount},
			metrics.GetMetricValueFromBigInt(vestAmount.BigInt()),
			[]gometrics.Label{metrics.GetLabelForStringValue(metrics.VesterAccount, entry.VesterAccount)},
		)
		// Report vester account balance after vest event.
		balanceAfterVest := k.bankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(entry.VesterAccount), entry.Denom)
		telemetry.SetGaugeWithLabels(
			[]string{types.ModuleName, metrics.BalanceAfterVestEvent},
			metrics.GetMetricValueFromBigInt(balanceAfterVest.Amount.BigInt()),
			[]gometrics.Label{metrics.GetLabelForStringValue(metrics.VesterAccount, entry.VesterAccount)},
		)
	}
}

func (k Keeper) GetAllVestEntries(ctx sdk.Context) (
	list []types.VestEntry,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VestEntryKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var val types.VestEntry
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return list
}

func (k Keeper) GetVestEntry(ctx sdk.Context, vesterAccount string) (
	val types.VestEntry,
	err error,
) {
	defer telemetry.ModuleMeasureSince(
		types.ModuleName,
		time.Now(),
		metrics.GetVestEntry,
		metrics.Latency,
	)

	// vest 계정에 해당하는 vest entry binary 가져오기
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VestEntryKeyPrefix))
	b := store.Get([]byte(vesterAccount))

	// If VestEntry does not exist in state, return error
	if b == nil {
		return types.VestEntry{}, errorsmod.Wrapf(types.ErrVestEntryNotFound, "vesterAccount: %s", vesterAccount)
	}

	// vest entry 역직렬화 후 리턴
	k.cdc.MustUnmarshal(b, &val)
	return val, nil
}

func (k Keeper) SetVestEntry(
	ctx sdk.Context,
	entry types.VestEntry,
) (
	err error,
) {
	// 인수가 정상적인이 검증
	if err := entry.Validate(); err != nil {
		return err
	}

	// KVStore에 vest 계정에 대한 vest entry 저장
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VestEntryKeyPrefix))
	b := k.cdc.MustMarshal(&entry)
	store.Set([]byte(entry.VesterAccount), b)
	return nil
}

func (k Keeper) DeleteVestEntry(
	ctx sdk.Context,
	vesterAccount string,
) (
	err error,
) {
	if _, err := k.GetVestEntry(ctx, vesterAccount); err != nil {
		return errorsmod.Wrap(err, "failed to delete vest entry")
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.VestEntryKeyPrefix))
	// KVStore에서 vest entry 삭제
	store.Delete([]byte(vesterAccount))

	return nil
}
