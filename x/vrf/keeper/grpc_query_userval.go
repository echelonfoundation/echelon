package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UservalAll(c context.Context, req *types.QueryAllUservalRequest) (*types.QueryAllUservalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var uservals []types.Userval
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	uservalStore := prefix.NewStore(store, types.KeyPrefix(types.UservalKeyPrefix))

	pageRes, err := query.Paginate(uservalStore, req.Pagination, func(key []byte, value []byte) error {
		var userval types.Userval
		if err := k.cdc.Unmarshal(value, &userval); err != nil {
			return err
		}

		uservals = append(uservals, userval)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUservalResponse{Userval: uservals, Pagination: pageRes}, nil
}

func (k Keeper) Userval(c context.Context, req *types.QueryGetUservalRequest) (*types.QueryGetUservalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetUserval(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetUservalResponse{Userval: val}, nil
}
