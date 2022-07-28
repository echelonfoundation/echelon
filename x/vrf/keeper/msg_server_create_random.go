package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

func (k msgServer) CreateRandom(goCtx context.Context, msg *types.MsgCreateRandom) (*types.MsgCreateRandomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CreateRandomNumber(ctx, msg)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateRandomResponse{}, err
}
