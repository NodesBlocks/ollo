package ons

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"ollo/x/ons/keeper"
	"ollo/x/ons/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the whois
	for _, elem := range genState.WhoisList {
		k.SetWhois(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WhoisList = k.GetAllWhois(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
