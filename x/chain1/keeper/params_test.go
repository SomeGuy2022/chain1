package keeper_test

import (
	"testing"

	testkeeper "github.com/SomeGuy2022/chain1/testutil/keeper"
	"github.com/SomeGuy2022/chain1/x/chain1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Chain1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
