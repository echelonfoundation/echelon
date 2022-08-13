package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/echelonfoundation/echelon/v3/testutil/keeper"
	"github.com/echelonfoundation/echelon/v3/x/vrf/keeper"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRandomval(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Randomval {
	items := make([]types.Randomval, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRandomval(ctx, items[i])
	}
	return items
}

func TestRandomvalGet(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNRandomval(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRandomval(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestRandomvalRemove(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNRandomval(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRandomval(ctx,
			item.Index,
		)
		_, found := keeper.GetRandomval(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRandomvalGetAll(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNRandomval(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllRandomval(ctx))
}
