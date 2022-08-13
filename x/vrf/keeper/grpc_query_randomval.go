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

func (k Keeper) RandomvalAll(c context.Context, req *types.QueryAllRandomvalRequest) (*types.QueryAllRandomvalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var randomvals []types.Randomval
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	randomvalStore := prefix.NewStore(store, types.KeyPrefix(types.RandomvalKeyPrefix))

	pageRes, err := query.Paginate(randomvalStore, req.Pagination, func(key []byte, value []byte) error {
		var randomval types.Randomval
		if err := k.cdc.Unmarshal(value, &randomval); err != nil {
			return err
		}

		randomvals = append(randomvals, randomval)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRandomvalResponse{Randomval: randomvals, Pagination: pageRes}, nil
}

func (k Keeper) Randomval(c context.Context, req *types.QueryGetRandomvalRequest) (*types.QueryGetRandomvalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetRandomval(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetRandomvalResponse{Randomval: val}, nil
}
