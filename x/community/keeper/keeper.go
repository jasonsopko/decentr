package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gofrs/uuid"

	"github.com/Decentr-net/decentr/x/community/types"
)

type TokenKeeper interface {
	AddTokens(ctx sdk.Context, owner sdk.AccAddress, amount sdk.Int)
}

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
	tokens   TokenKeeper
}

// NewKeeper creates new instances of the community Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, tokens TokenKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		tokens:   tokens,
	}
}

// CreatePost creates new post. Keeper's key is joined owner and uuid.
func (k Keeper) CreatePost(ctx sdk.Context, owner sdk.AccAddress, p types.Post) {
	store := ctx.KVStore(k.storeKey)

	key := append(owner.Bytes(), p.UUID[:]...)

	store.Set(key, k.cdc.MustMarshalBinaryBare(p))

	// store to index
}

// DeletePost deletes the post from keeper.
func (k Keeper) DeletePost(ctx sdk.Context, owner sdk.AccAddress, id uuid.UUID) {
	store := ctx.KVStore(k.storeKey)

	key := append(owner.Bytes(), id.Bytes()...)

	store.Delete(key)

	// remove from index
}

// GetPostByKey returns entire post by keeper's key.
func (k Keeper) GetPostByKey(ctx sdk.Context, key []byte) types.Post {
	store := ctx.KVStore(k.storeKey)

	if !store.Has(key) {
		return types.Post{}
	}

	bz := store.Get(key)

	var post types.Post
	k.cdc.MustUnmarshalBinaryBare(bz, &post)
	return post
}

// GetPostsIterator returns an iterator over all posts
func (k Keeper) GetPostsIterator(ctx sdk.Context, prefix string) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(prefix))
}