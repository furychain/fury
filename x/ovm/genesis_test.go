package ovm_test

import (
	"testing"

	"github.com/merlin-network/merlin/testutil/nullify"
	simappUtil "github.com/merlin-network/merlin/testutil/simapp"
	"github.com/merlin-network/merlin/x/ovm"
	"github.com/merlin-network/merlin/x/ovm/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		KeyVault: types.KeyVault{
			PublicKeys: []string{"Key1"},
		},
	}

	tApp, ctx, err := simappUtil.GetTestObjects()
	require.NoError(t, err)

	ovm.InitGenesis(ctx, tApp.OVMKeeper, genesisState)
	got := ovm.ExportGenesis(ctx, tApp.OVMKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
