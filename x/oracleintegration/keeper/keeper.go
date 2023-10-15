package keeper

import (
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"

	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		scopedKeeper  exported.ScopedKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper exported.ScopedKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SaveRequest(ctx sdk.Context, requestID uint64, request types.RequestPrice) {
	ctx.KVStore(k.storeKey).Set(types.KeyRequestIdPrefix(requestID), k.cdc.MustMarshal(&request))
}

func (k Keeper) GetRequest(ctx sdk.Context, requestID uint64) (types.RequestPrice, error) {
	var priceRequest types.RequestPrice
	request := ctx.KVStore(k.storeKey).Get(types.KeyRequestIdPrefix(requestID))
	if len(request) == 0 {
		return priceRequest, sdkerrors.Wrapf(types.ErrRequestIdNotFound, "price request for request %d not found", requestID)
	}

	err := k.cdc.Unmarshal(request, &priceRequest)

	return priceRequest, err
}

func (k Keeper) SaveResult(ctx sdk.Context, requestID uint64, fetchResult types.OracleRequestResult) {
	ctx.KVStore(k.storeKey).Set(types.KeyRequestIdResultPrefix(requestID), k.cdc.MustMarshal(&fetchResult))
}

func (k Keeper) GetResult(ctx sdk.Context, requestID uint64) (types.OracleRequestResult, error) {
	var oracleResult types.OracleRequestResult
	result := ctx.KVStore(k.storeKey).Get(types.KeyRequestIdResultPrefix(requestID))
	if len(result) == 0 {
		return oracleResult, sdkerrors.Wrapf(types.ErrResultNotFound, "result for request %d not found", requestID)
	}

	err := k.cdc.Unmarshal(result, &oracleResult)

	return oracleResult, err
}

func (k Keeper) GetResults(ctx sdk.Context) ([]types.ResultResponse, error) {
	var results []types.ResultResponse
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyRequestIdResult)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oracleResult types.OracleRequestResult
		err := k.cdc.Unmarshal(iterator.Value(), &oracleResult)
		if err != nil {
			return nil, err
		}

		key := iterator.Key()[len(types.KeyRequestIdResult):]
		results = append(results, types.ResultResponse{RequestId: sdk.BigEndianToUint64(key), Result: oracleResult})
	}

	return results, nil
}

func (k Keeper) GetRequests(ctx sdk.Context) ([]types.RequestResponse, error) {
	var requests []types.RequestResponse
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyRequestId)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var priceRequest types.RequestPrice
		err := k.cdc.Unmarshal(iterator.Value(), &priceRequest)
		if err != nil {
			return nil, err
		}

		key := iterator.Key()[len(types.KeyRequestId):]
		requests = append(requests, types.RequestResponse{
			RequestId: sdk.BigEndianToUint64(key),
			Request:   priceRequest,
		})
	}

	return requests, nil
}
