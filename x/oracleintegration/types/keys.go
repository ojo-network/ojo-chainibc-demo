package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "oracleintegration"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracleintegration"

	FetchPriceClientIDKey = "fetch_prices"

	// Version defines the current version the IBC module supports.
	Version = "relayoracle-1"

	// PortID is the default port id that module binds to.
	PortID = "relayoracle"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey            = KeyPrefix("oracleintegration-port-")
	KeyRequestId       = KeyPrefix("request-id")
	KeyRequestIdResult = KeyPrefix("result-id")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeyRequestIdPrefix(requestID uint64) (key []byte) {
	key = append(key, KeyRequestId...)
	return append(key, sdk.Uint64ToBigEndian(requestID)...)
}

func KeyRequestIdResultPrefix(requestID uint64) (key []byte) {
	key = append(key, KeyRequestIdResult...)
	return append(key, sdk.Uint64ToBigEndian(requestID)...)
}
