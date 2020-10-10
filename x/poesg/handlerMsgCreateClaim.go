package poesg

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukitsbrian/poe-sg/x/poesg/types"
	"github.com/lukitsbrian/poe-sg/x/poesg/keeper"
)

func handleMsgCreateClaim(ctx sdk.Context, k keeper.Keeper, claim *types.MsgClaim) (*sdk.Result, error) {
	k.CreateClaim(ctx, *claim)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
