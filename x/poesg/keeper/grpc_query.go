package keeper

import (
	"github.com/lukitsbrian/poe-sg/x/poesg/types"
)

var _ types.QueryServer = Keeper{}
