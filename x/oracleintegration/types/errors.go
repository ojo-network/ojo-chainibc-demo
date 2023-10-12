package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/oracleintegration module sentinel errors
var (
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrRequestingPrices = sdkerrors.Register(ModuleName, 1503, "error requesting prices")
)
