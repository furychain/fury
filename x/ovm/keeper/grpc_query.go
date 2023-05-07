package keeper

import (
	"github.com/merlin-network/merlin/x/ovm/types"
)

var _ types.QueryServer = Keeper{}
