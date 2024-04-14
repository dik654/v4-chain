package keeper

import (
	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dydxprotocol/v4-chain/protocol/dtypes"
	"github.com/dydxprotocol/v4-chain/protocol/x/vault/types"
)

// GetTotalShares gets TotalShares for a vault.
func (k Keeper) GetTotalShares(
	ctx sdk.Context,
	vaultId types.VaultId,
) (val types.NumShares, exists bool) {
	// KVStore 생성
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TotalSharesKeyPrefix))

	// ToStateKey는 ID를 bytes로 직렬화
	// store에서 ID를 key로하여 share 얻기
	b := store.Get(vaultId.ToStateKey())
	// share가 nil이라면 에러
	if b == nil {
		return val, false
	}

	// 직렬화되어있는 share를 역직렬화
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetTotalShares sets TotalShares for a vault. Returns error if `totalShares` is negative.
func (k Keeper) SetTotalShares(
	ctx sdk.Context,
	vaultId types.VaultId,
	totalShares types.NumShares,
) error {
	// Cmp는 "github.com/dydxprotocol/v4-chain/protocol/lib"에서 선언
	// share의 개수가 0보다 작다면 err
	if totalShares.NumShares.Cmp(dtypes.NewInt(0)) == -1 {
		return types.ErrNegativeShares
	}

	// totalShare를 totalShare KVStore에 저장
	b := k.cdc.MustMarshal(&totalShares)
	totalSharesStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TotalSharesKeyPrefix))
	totalSharesStore.Set(vaultId.ToStateKey(), b)

	return nil
}
