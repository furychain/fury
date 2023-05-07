package keeper

import (
	"github.com/merlin-network/merlin/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
