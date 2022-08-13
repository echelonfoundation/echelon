package vrf

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/keeper"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the randomval
	for _, elem := range genState.RandomvalList {
		k.SetRandomval(ctx, elem)
	}
	// Set all the userval
	for _, elem := range genState.UservalList {
		k.SetUserval(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.RandomvalList = k.GetAllRandomval(ctx)
	genesis.UservalList = k.GetAllUserval(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
