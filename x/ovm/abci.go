package ovm

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlin-network/merlin/x/ovm/keeper"
)

// EndBlocker settles the active bets of resolved markets
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	err := k.FinishProposals(ctx)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("end block number %d error: %s", ctx.BlockHeight(), err))
	}
}