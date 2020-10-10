package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lukitsbrian/poe-sg/x/poesg/types"
)

func (k Keeper) CreateClaim(ctx sdk.Context, claim types.MsgClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	b := k.cdc.MustMarshalBinaryBare(&claim)
	store.Set(types.KeyPrefix(types.ClaimKey), b)
}

// GetClaim returns the claim information
func (k Keeper) GetClaim(ctx sdk.Context, key string) (types.MsgClaim, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	var claim types.MsgClaim
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &claim)
	if err != nil {
		return claim, err
	}
	return claim, nil
}

// SetClaim sets a claim
func (k Keeper) SetClaim(ctx sdk.Context, claim types.MsgClaim) {
	claimKey := claim.Proof
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&claim)
	key := []byte(claimKey)
	store.Set(key, bz)
}

func (k Keeper) DeleteClaim(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	store.Delete([]byte(key))
}

func (k Keeper) GetAllClaim(ctx sdk.Context) (msgs []types.MsgClaim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ClaimKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgClaim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// Get creator of the item
func (k Keeper) GetClaimOwner(ctx sdk.Context, key string) sdk.AccAddress {
	claim, err := k.GetClaim(ctx, key)
	if err != nil {
		return nil
	}
	return claim.Creator
}

// Check if the key exists in the store
func (k Keeper) ClaimExists(ctx sdk.Context, key string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	return store.Has([]byte(key))
}
