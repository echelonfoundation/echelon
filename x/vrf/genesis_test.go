package vrf_test

import (
	"testing"

	keepertest "github.com/echelonfoundation/echelon/v3/testutil/keeper"
	"github.com/echelonfoundation/echelon/v3/x/vrf"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		RandomvalList: []types.Randomval{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		UservalList: []types.Userval{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VRFKeeper(t)
	random.InitGenesis(ctx, *k, genesisState)
	got := random.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.RandomvalList, len(genesisState.RandomvalList))
	require.Subset(t, genesisState.RandomvalList, got.RandomvalList)
	require.Len(t, got.UservalList, len(genesisState.UservalList))
	require.Subset(t, genesisState.UservalList, got.UservalList)
	// this line is used by starport scaffolding # genesis/test/assert
}
