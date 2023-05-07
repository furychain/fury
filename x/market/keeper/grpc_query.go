package keeper

import (
	"github.com/furychain/fury/x/market/types"
)

var _ types.QueryServer = Keeper{}
