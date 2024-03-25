package keeper

import (
	"math/big"
	"sort"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	indexerevents "github.com/dydxprotocol/v4-chain/protocol/indexer/events"
	"github.com/dydxprotocol/v4-chain/protocol/indexer/indexer_manager"
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/x/assets/types"
)

func (k Keeper) CreateAsset(
	ctx sdk.Context,
	assetId uint32,
	symbol string,
	denom string,
	denomExponent int32,
	hasMarket bool,
	marketId uint32,
	atomicResolution int32,
) (types.Asset, error) {
	// 등록하려는 assetId에 해당하는 asset이 이미 존재한다면
	if prevAsset, exists := k.GetAsset(ctx, assetId); exists {
		// asset이 이미 존재한다는 에러 리턴
		return types.Asset{}, errorsmod.Wrapf(
			types.ErrAssetIdAlreadyExists,
			"previous asset = %v",
			prevAsset,
		)
	}

	// assetId가 지정된 USDC의 asset id라면
	if assetId == types.AssetUsdc.Id {
		// Ensure assetId zero is always USDC. This is a protocol-wide invariant.
		// 지정한 denom 식별자(ERC20에서의 name과 유사)가 USDC와 동일하다면
		if denom != types.AssetUsdc.Denom {
			return types.Asset{}, types.ErrUsdcMustBeAssetZero
		}

		// Confirm that USDC asset has the expected denom exponent (-6).
		// This is an important invariant before coin-to-quote-quantum conversion
		// is correctly implemented. See CLOB-871 for details.
		// 지정한 단위 exponent(ex. ether, wei)가 USDC의 단위 exponent와 동일하지 않다면 에러
		if denomExponent != types.AssetUsdc.DenomExponent {
			return types.Asset{}, errorsmod.Wrapf(
				types.ErrUnexpectedUsdcDenomExponent,
				"expected = %v, actual = %v",
				types.AssetUsdc.DenomExponent,
				denomExponent,
			)
		}
	}

	// Ensure USDC is not created with a non-zero assetId. This is a protocol-wide invariant.
	// USDC의 assetId가 아니지만, USDC의 식별자를 사용하고 있다면 에러
	if assetId != types.AssetUsdc.Id && denom == types.AssetUsdc.Denom {
		return types.Asset{}, types.ErrUsdcMustBeAssetZero
	}

	// Ensure the denom is unique versus existing assets.
	// 현재 존재하는 모든 asset을 가져와서
	allAssets := k.GetAllAssets(ctx)
	for _, asset := range allAssets {
		// 추가하려는 asset의 denom 식별자가 현존하는 asset과 중복되는 경우 error
		if asset.Denom == denom {
			return types.Asset{}, errorsmod.Wrap(types.ErrAssetDenomAlreadyExists, denom)
		}
	}

	// Create the asset
	// 위의 예외 사항에 결격사유가 없다면 인수에 맞춰 asset 구조체 생성
	asset := types.Asset{
		Id:               assetId,
		Symbol:           symbol,
		Denom:            denom,
		DenomExponent:    denomExponent,
		HasMarket:        hasMarket,
		MarketId:         marketId,
		AtomicResolution: atomicResolution,
	}

	// Validate market
	// 마켓이 존재한다고 설정했을 경우
	if hasMarket {
		// 지정한 marketId의 마켓 가격 가져와서
		// 실패할 경우 에러 리턴
		if _, err := k.pricesKeeper.GetMarketPrice(ctx, marketId); err != nil {
			return asset, err
		}
	// 마켓이 없다고 설정했지만 marketId가 지정되어있다면 에러 리턴
	} else if marketId > 0 {
		return asset, errorsmod.Wrapf(
			types.ErrInvalidMarketId,
			"Market ID: %v",
			marketId,
		)
	}

	// Store the new asset
	// market에 대한 문제가 해결됐다면 asset 구조체를 keeper에 저장
	k.setAsset(ctx, asset)

	// 이벤트 생성
	k.GetIndexerEventManager().AddTxnEvent(
		ctx,
		indexerevents.SubtypeAsset,
		indexerevents.AssetEventVersion,
		indexer_manager.GetBytes(
			indexerevents.NewAssetCreateEvent(
				assetId,
				asset.Symbol,
				asset.HasMarket,
				asset.MarketId,
				asset.AtomicResolution,
			),
		),
	)

	return asset, nil
}

func (k Keeper) ModifyAsset(
	ctx sdk.Context,
	id uint32,
	hasMarket bool,
	marketId uint32,
) (types.Asset, error) {
	// Get asset
	// assetId에 해당하는 asset이 존재하는지 체크
	asset, exists := k.GetAsset(ctx, id)
	// 없다면 에러
	if !exists {
		return asset, errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
	}

	// Validate market
	// 변경하려는 market의 가격을 가져오고
	// 가격 가져오기에 실패할 경우 에러 리턴
	if _, err := k.pricesKeeper.GetMarketPrice(ctx, marketId); err != nil {
		return asset, err
	}

	// Modify asset
	// 문제가 없다면 인수에 넣은대로 수정
	asset.HasMarket = hasMarket
	asset.MarketId = marketId

	// Store the modified asset
	// 수정한 asset 구조체 keeper에 저장
	k.setAsset(ctx, asset)

	return asset, nil
}

func (k Keeper) setAsset(
	ctx sdk.Context,
	asset types.Asset,
) {
	// asset 구조체 binary 마샬링
	b := k.cdc.MustMarshal(&asset)
	// keeper의 KVStore에 저장
	assetStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AssetKeyPrefix))
	assetStore.Set(lib.Uint32ToKey(asset.Id), b)
}

func (k Keeper) GetAsset(
	ctx sdk.Context,
	id uint32,
) (val types.Asset, exists bool) {
	// keeper의 KVStore 가져오기
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AssetKeyPrefix))

	// KVStore에서 id에 해당하는 binary 마샬링된 asset 구조체 가져오기
	b := store.Get(lib.Uint32ToKey(id))
	if b == nil {
		return val, false
	}

	// 구조체로 언마샬링
	k.cdc.MustUnmarshal(b, &val)
	// 구조체 리턴
	return val, true
}

func (k Keeper) GetAllAssets(
	ctx sdk.Context,
) (list []types.Asset) {
	// keeper의 KVStore 가져오기
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AssetKeyPrefix))
	// KVStore용 iterator 가져오기
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	// KVStore에 존재하는 모든 binary 상태의 asset 구조체를 돌면서
	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		// 구조체로 언마샬링 후
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		// 리턴 할 asset 슬라이스에 언마샬링된 구조체 추가
		list = append(list, val)
	}

	// 슬라이스 내 구조체들을 assetId 오름차순으로 정렬 
	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})

	return list
}

// GetNetCollateral returns the net collateral that a given position (quantums)
// for a given assetId contributes to an account.
func (k Keeper) GetNetCollateral(
	ctx sdk.Context,
	id uint32,
	bigQuantums *big.Int,
) (
	bigNetCollateralQuoteQuantums *big.Int,
	err error,
) {
	if id == types.AssetUsdc.Id {
		return new(big.Int).Set(bigQuantums), nil
	}

	// Get asset
	_, exists := k.GetAsset(ctx, id)
	if !exists {
		return big.NewInt(0), errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
	}

	// Balance is zero.
	if bigQuantums.BitLen() == 0 {
		return big.NewInt(0), nil
	}

	// Balance is positive.
	// TODO(DEC-581): add multi-collateral support.
	if bigQuantums.Sign() == 1 {
		return big.NewInt(0), types.ErrNotImplementedMulticollateral
	}

	// Balance is negative.
	// TODO(DEC-582): add margin-trading support.
	return big.NewInt(0), types.ErrNotImplementedMargin
}

// GetMarginRequirements returns the initial and maintenance margin-
// requirements for a given position size for a given assetId.
func (k Keeper) GetMarginRequirements(
	ctx sdk.Context,
	id uint32,
	bigQuantums *big.Int,
) (
	bigInitialMarginQuoteQuantums *big.Int,
	bigMaintenanceMarginQuoteQuantums *big.Int,
	err error,
) {
	// QuoteBalance does not contribute to any margin requirements.
	if id == types.AssetUsdc.Id {
		return big.NewInt(0), big.NewInt(0), nil
	}

	// Get asset
	_, exists := k.GetAsset(ctx, id)
	if !exists {
		return big.NewInt(0), big.NewInt(0), errorsmod.Wrap(
			types.ErrAssetDoesNotExist, lib.UintToString(id))
	}

	// Balance is zero or positive.
	if bigQuantums.Sign() >= 0 {
		return big.NewInt(0), big.NewInt(0), nil
	}

	// Balance is negative.
	// TODO(DEC-582): margin-trading
	return big.NewInt(0), big.NewInt(0), types.ErrNotImplementedMargin
}

// ConvertAssetToCoin converts the given `assetId` and `quantums` used in `x/asset`,
// to an `sdk.Coin` in correspoding `denom` and `amount` used in `x/bank`.
// Also outputs `convertedQuantums` which has the equal value as converted `sdk.Coin`.
// The conversion is done with the formula:
//
//	denom_amount = quantums * 10^(atomic_resolution - denom_exponent)
//
// If the resulting `denom_amount` is not an integer, it is rounded down,
// and `convertedQuantums` of the equal value is returned. The upstream
// transfer function should adjust asset balance with `convertedQuantums`
// to ensure that that no fund is ever lost in the conversion due to rounding error.
//
// Example:
// Assume `denom_exponent` = -7, `atomic_resolution` = -8,
// ConvertAssetToCoin(`101 quantums`) should output:
// - `convertedQuantums` = 100 quantums
// -  converted coin amount = 10 coin
func (k Keeper) ConvertAssetToCoin(
	ctx sdk.Context,
	assetId uint32,
	quantums *big.Int,
) (
	convertedQuantums *big.Int,
	coin sdk.Coin,
	err error,
) {
	asset, exists := k.GetAsset(ctx, assetId)
	if !exists {
		return nil, sdk.Coin{}, errorsmod.Wrap(
			types.ErrAssetDoesNotExist, lib.UintToString(assetId))
	}

	if lib.AbsInt32(asset.AtomicResolution) > types.MaxAssetUnitExponentAbs {
		return nil, sdk.Coin{}, errorsmod.Wrapf(
			types.ErrInvalidAssetAtomicResolution,
			"asset: %+v",
			asset,
		)
	}

	if lib.AbsInt32(asset.DenomExponent) > types.MaxAssetUnitExponentAbs {
		return nil, sdk.Coin{}, errorsmod.Wrapf(
			types.ErrInvalidDenomExponent,
			"asset: %+v",
			asset,
		)
	}

	bigRatDenomAmount := lib.BigMulPow10(
		quantums,
		asset.AtomicResolution-asset.DenomExponent,
	)

	// round down to get denom amount that was converted.
	bigConvertedDenomAmount := lib.BigRatRound(bigRatDenomAmount, false)

	bigRatConvertedQuantums := lib.BigMulPow10(
		bigConvertedDenomAmount,
		asset.DenomExponent-asset.AtomicResolution,
	)

	bigConvertedQuantums := bigRatConvertedQuantums.Num()

	return bigConvertedQuantums, sdk.NewCoin(asset.Denom, sdkmath.NewIntFromBigInt(bigConvertedDenomAmount)), nil
}

// IsPositionUpdatable returns whether position of an asset is updatable.
func (k Keeper) IsPositionUpdatable(
	ctx sdk.Context,
	id uint32,
) (
	updatable bool,
	err error,
) {
	_, exists := k.GetAsset(ctx, id)
	if !exists {
		return false, errorsmod.Wrap(types.ErrAssetDoesNotExist, lib.UintToString(id))
	}
	// All existing assets are by default updatable.
	return true, nil
}
