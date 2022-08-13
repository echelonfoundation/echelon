package keeper

import (
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

var _ types.QueryServer = Keeper{}
