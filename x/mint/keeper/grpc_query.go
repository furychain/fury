package keeper

import (
	"github.com/furychain/fury/x/mint/types"
)

var _ types.QueryServer = Keeper{}
