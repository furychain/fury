package keeper

import (
	"github.com/furychain/fury/x/bet/types"
)

var _ types.QueryServer = Keeper{}
