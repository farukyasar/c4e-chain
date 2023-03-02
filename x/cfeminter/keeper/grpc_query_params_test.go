package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/chain4energy/c4e-chain/testutil/keeper"
	testminter "github.com/chain4energy/c4e-chain/testutil/module/cfeminter"
	"github.com/chain4energy/c4e-chain/x/cfeminter/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx, _ := testkeeper.CfeminterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	params.MintDenom = "denom"
	params.Minters = createLinearMintings(time.Now())

	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	testminter.CompareCfeminterParams(t, params, response.Params)
}
