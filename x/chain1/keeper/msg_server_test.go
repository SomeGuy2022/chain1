package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/SomeGuy2022/chain1/testutil/keeper"
	"github.com/SomeGuy2022/chain1/x/chain1/keeper"
	"github.com/SomeGuy2022/chain1/x/chain1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Chain1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
