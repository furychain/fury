package keeper

import (
	"github.com/merlin-network/merlin/x/market/types"
)

var _ types.QueryServer = Keeper{}
