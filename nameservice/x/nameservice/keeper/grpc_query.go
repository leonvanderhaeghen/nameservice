package keeper

import (
	"github.com/leonvanderhaeghen/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
