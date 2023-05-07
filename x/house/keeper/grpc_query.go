package keeper

import (
	"github.com/merlin-network/merlin/x/house/types"
)

var _ types.QueryServer = Keeper{}
