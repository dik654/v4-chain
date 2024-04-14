package keeper

import (
	"context"

	"github.com/dydxprotocol/v4-chain/protocol/x/vault/types"
)

// DepositToVault deposits from a subaccount to a vault.
// subaccount에서 vault로 deposit
func (k msgServer) DepositToVault(
	goCtx context.Context,
	msg *types.MsgDepositToVault,
) (*types.MsgDepositToVaultResponse, error) {
	return &types.MsgDepositToVaultResponse{}, nil
}
