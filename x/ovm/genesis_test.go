package ovm_test

import (
	"testing"

	"github.com/furychain/fury/testutil/nullify"
	simappUtil "github.com/furychain/fury/testutil/simapp"
	"github.com/furychain/fury/x/ovm"
	"github.com/furychain/fury/x/ovm/types"
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
