package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	// storeKey type 가져오기용
	storetypes "cosmossdk.io/store/types"
	// 직렬화
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// dydx용 라이브러리
	"github.com/dydxprotocol/v4-chain/protocol/lib"
	"github.com/dydxprotocol/v4-chain/protocol/x/vault/types"
)

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
		cdc:      cdc,
		storeKey: storeKey,
		// lib/collections.go
		// slice는 KV mapping으로 변환
		// key만 authority로 지정. value는 struct{}{}
		authorities: lib.UniqueSliceToSet(authorities),
	}
}

// key가 존재하는지 확인
func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

// logger 리턴
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {}
