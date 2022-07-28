package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VerifyValues(goCtx context.Context, req *types.QueryVerifyValuesRequest) (*types.QueryVerifyValuesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	is_verified, err := k.VerifyRandomNumber(ctx, req)

	// TODO: Process the query
	_ = ctx

	return &types.QueryVerifyValuesResponse{Verified: is_verified}, err
}
