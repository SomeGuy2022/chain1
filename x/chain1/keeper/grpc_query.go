package keeper

import (
	"github.com/SomeGuy2022/chain1/x/chain1/types"
)

var _ types.QueryServer = Keeper{}
