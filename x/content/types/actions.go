package types

import (
	"github.com/bitsongofficial/go-bitsong/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Actions map[string]map[int64]sdk.Coin

func NewActions() Actions {
	zero := sdk.NewCoin(types.BondDenom, sdk.ZeroInt())
	return Actions{
		"promotions": {
			0: zero,
		},
		"streams": {
			0: zero,
		},
		"downloads": {
			0: zero,
		},
		"tips": {
			0: zero,
		},
		"vote": {
			0: zero, // +1 or -1
		},
	}
}