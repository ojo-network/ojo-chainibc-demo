package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) RequestPrice(goCtx context.Context, request *types.MsgRequestPrice) (*types.MsgRequestPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	requestedSymbols := request.GetSymbols()
	params := m.GetParams(ctx)
	channelCap, ok := m.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(types.PortID, params.GetSourceChannel()))
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrRequestingPrices, "error requesting %v", types.PriceRequestType_name[int32(request.Type)])
	}

	oracleRequest := types.OracleRequestPacketData{
		ClientID: types.FetchPriceClientIDKey,
		Calldata: m.cdc.MustMarshal(&types.RequestPrice{
			Denoms:  requestedSymbols,
			Request: request.GetType(),
		}),
	}

	_, err := m.channelKeeper.SendPacket(
		ctx,
		channelCap,
		types.PortID,
		params.GetSourceChannel(),
		clienttypes.NewHeight(0, 0),
		uint64(ctx.BlockTime().UnixNano()+int64(20*time.Minute)),
		types.ModuleCdc.MustMarshal(&oracleRequest),
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgRequestPriceResponse{}, nil
}
