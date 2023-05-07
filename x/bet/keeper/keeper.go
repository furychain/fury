package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/furychain/fury/x/bet/types"
)

// Keeper is the type for module properties
type Keeper struct {
	cdc          codec.BinaryCodec
	storeKey     sdk.StoreKey
	memKey       sdk.StoreKey
	paramstore   paramtypes.Subspace
	marketKeeper types.MarketKeeper
	srKeeper     types.SRKeeper
	ovmKeeper    types.OVMKeeper
}

// NewKeeper creates new keeper object
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

// SetMarketKeeper sets market keeper to the bet keeper.
func (k *Keeper) SetMarketKeeper(marketKeeper types.MarketKeeper) {
	k.marketKeeper = marketKeeper
}

// SetSRKeeper sets sr keeper to the bet keeper.
func (k *Keeper) SetSRKeeper(srKeeper types.SRKeeper) {
	k.srKeeper = srKeeper
}

// SetOVMKeeper sets ovm keeper to the bet keeper.
func (k *Keeper) SetOVMKeeper(ovmKeeper types.OVMKeeper) {
	k.ovmKeeper = ovmKeeper
}

// Logger returns the logger of the keeper
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
