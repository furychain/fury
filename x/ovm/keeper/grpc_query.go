package keeper

import (
	"github.com/furychain/fury/x/ovm/types"
)

var _ types.QueryServer = Keeper{}
