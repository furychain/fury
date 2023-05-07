package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlin-network/merlin/cmd/merlind/cmd"
	"github.com/merlin-network/merlin/testutil/network"
	"github.com/stretchr/testify/require"
)

func TestAddGenesisAccountCmdPanic(t *testing.T) {
	userHomeDir, err := os.UserHomeDir()
	require.NoError(t, err)
	defaultNodeHome := filepath.Join(userHomeDir, ".merlin")
	panicFunc := func() {
		cmd.AddGenesisAccountCmd(defaultNodeHome)
	}

	require.NotPanics(t, panicFunc)
}

func TestSampleCMD(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx
	sec := hd.Secp256k1

	keyInfo, _, err := ctx.Keyring.NewMnemonic("genUser1", keyring.English, sdk.FullFundraiserPath, "", sec)
	require.NoError(t, err)

	userHomeDir, err := os.UserHomeDir()
	require.NoError(t, err)
	defaultNodeHome := filepath.Join(userHomeDir, ".merlin")

	fields := []string{keyInfo.GetAddress().String(), "10000000umerlin"}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		code uint32
	}{
		{
			desc: "valid",
			args: []string{},
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{}
			args = append(args, fields...)
			args = append(args, tc.args...)
			_, err := clitestutil.ExecTestCLICmd(ctx, cmd.AddGenesisAccountCmd(defaultNodeHome), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
