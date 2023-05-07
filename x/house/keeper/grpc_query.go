package keeper

import (
	"github.com/furychain/fury/x/house/types"
)

var _ types.QueryServer = Keeper{}
