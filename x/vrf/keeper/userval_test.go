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

func createNUserval(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Userval {
	items := make([]types.Userval, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserval(ctx, items[i])
	}
	return items
}

func TestUservalGet(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNUserval(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserval(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestUservalRemove(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNUserval(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserval(ctx,
			item.Index,
		)
		_, found := keeper.GetUserval(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUservalGetAll(t *testing.T) {
	keeper, ctx := keepertest.VRFKeeper(t)
	items := createNUserval(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllUserval(ctx))
}
