package keeper

import (
	"github.com/Fairblock/fairyring/x/pep/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	coin := k.MinGasPrice(ctx)
	return types.NewParams(
		k.TrustedAddresses(ctx),
		k.TrustedCounterParties(ctx),
		k.KeyshareChannelID(ctx),
		&coin,
		k.IsSourceChain(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.Set(ctx, params)
}

// TrustedAddresses returns the TrustedAddresses param
func (k Keeper) TrustedAddresses(ctx sdk.Context) (res []string) {
	store, _ := k.params.Get(ctx)
	res = store.GetTrustedAddresses()
	return
}

// TrustedCounterParties returns the TrustedCounterParties param
func (k Keeper) TrustedCounterParties(ctx sdk.Context) (res []*types.TrustedCounterParty) {
	store, _ := k.params.Get(ctx)
	res = store.GetTrustedCounterParties()
	return
}

// KeyshareChannelID returns the KeyshareChannelID param
func (k Keeper) KeyshareChannelID(ctx sdk.Context) (res string) {
	store, _ := k.params.Get(ctx)
	res = store.GetKeyshareChannelId()
	return
}

func (k Keeper) MinGasPrice(ctx sdk.Context) (res sdk.Coin) {
	store, _ := k.params.Get(ctx)
	res = *store.GetMinGasPrice()
	return
}

func (k Keeper) IsSourceChain(ctx sdk.Context) (res bool) {
	store, _ := k.params.Get(ctx)
	res = store.GetIsSourceChain()
	return
}
