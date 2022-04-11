package chain1_test

import (
	"testing"

	keepertest "github.com/SomeGuy2022/chain1/testutil/keeper"
	"github.com/SomeGuy2022/chain1/testutil/nullify"
	"github.com/SomeGuy2022/chain1/x/chain1"
	"github.com/SomeGuy2022/chain1/x/chain1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Chain1Keeper(t)
	chain1.InitGenesis(ctx, *k, genesisState)
	got := chain1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
