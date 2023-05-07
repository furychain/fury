package keeper

import (
	"github.com/furychain/fury/x/strategicreserve/types"
)

var _ types.QueryServer = Keeper{}
