package keeper

import (
	"github.com/merlin-network/merlin/x/mint/types"
)

var _ types.QueryServer = Keeper{}
