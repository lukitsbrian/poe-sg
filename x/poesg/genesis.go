package poesg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukitsbrian/poe-sg/x/poesg/keeper"
	"github.com/lukitsbrian/poe-sg/x/poesg/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return types.DefaultGenesis()
}
