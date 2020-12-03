package server

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/util/storehelpers"
	"github.com/regen-network/regen-ledger/x/bank"
	"github.com/regen-network/regen-ledger/x/bank/math"
)

func (s serverImpl) Balance(goCtx context.Context, req *bank.QueryBalanceRequest) (*bank.QueryBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	denom := req.Denom

	balance, err := storehelpers.GetDecimal(store, BalanceKey(addr, denom))
	if err != nil {
		return nil, err
	}

	return &bank.QueryBalanceResponse{Balance: &bank.Coin{
		Denom:  denom,
		Amount: math.DecimalString(balance),
	}}, nil
}

func (s serverImpl) SupplyOf(goCtx context.Context, req *bank.QuerySupplyOfRequest) (*bank.QuerySupplyOfResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)

	denom := req.Denom

	supply, err := storehelpers.GetDecimal(store, SupplyKey(denom))
	if err != nil {
		return nil, err
	}

	return &bank.QuerySupplyOfResponse{Amount: &bank.Coin{
		Denom:  denom,
		Amount: math.DecimalString(supply),
	}}, nil
}

func (s serverImpl) Precision(goCtx context.Context, request *bank.QueryPrecisionRequest) (*bank.QueryPrecisionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(s.key)
	x, err := storehelpers.GetUint32(store, MaxDecimalPlacesKey(request.Denom))
	if err != nil {
		return nil, err
	}

	return &bank.QueryPrecisionResponse{MaxDecimalPlaces: x}, nil
}
