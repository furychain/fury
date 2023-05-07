package keeper

import (
	"github.com/merlin-network/merlin/x/bet/types"
)

var _ types.QueryServer = Keeper{}
