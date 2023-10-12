package oracleintegration

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerror "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

func (im IBCModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData types.OracleResponsePacketData
	params := im.keeper.GetParams(ctx)
	if modulePacket.DestinationChannel != params.SourceChannel {
		return channeltypes.Acknowledgement{}, sdkerrors.Wrapf(sdkerror.ErrInvalidRequest, "packet destination and source does not match")
	}

	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}
	switch modulePacketData.GetClientID() {
	case types.FetchPriceClientIDKey:
		//TODO: remove
		ctx.Logger().Info("Received oracle packet ", "packet", modulePacketData.String())

		var priceResult types.OracleRequestResult
		if err := types.ModuleCdc.Unmarshal(modulePacketData.GetResult(), &priceResult); err != nil {
			ack = channeltypes.NewErrorAcknowledgement(err)
			return ack, sdkerrors.Wrap(sdkerror.ErrUnknownRequest,
				"cannot unmarshall price result packet")
		}

		//TODO: remove
		ctx.Logger().Info("price result oracle packet", "price result", priceResult.String())

		im.keeper.SaveFetchResult(ctx, modulePacketData.GetRequestID(), priceResult)
		//TODO: store the data here

	default:
		err := sdkerrors.Wrapf(sdkerror.ErrJSONUnmarshal,
			"market received packet not found: %s", modulePacketData.GetClientID())
		ack = channeltypes.NewErrorAcknowledgement(err)
		return ack, err
	}

	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			&types.OracleRequestPacketAcknowledgement{RequestID: modulePacketData.RequestID},
		),
	)
	return ack, nil
}

func (im IBCModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck types.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		//TODO :remove after testing
		ctx.Logger().Error("processing acknow response packet", "packet", oracleAck.String())

		var requestType protobuftypes.Int32Value
		err = types.ModuleCdc.UnmarshalLengthPrefixed(oracleAck.GetData(), &requestType)
		if err != nil {
			return nil, err
		}

		im.keeper.SaveRequestID(ctx, oracleAck.RequestID, requestType)
		return &sdk.Result{}, nil

		//TODO: fix after testing
		//default:
		//	return nil, sdkerrors.Wrapf(sdkerror.ErrJSONUnmarshal,
		//		"market acknowledgment packet not found: %s", orac.)
	}

	return nil, nil
}
