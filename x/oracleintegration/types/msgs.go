package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgRequestPrice{}
)

func NewMsgPriceRequest(symbols []string, requestType PriceRequestType, from string) *MsgRequestPrice {
	return &MsgRequestPrice{
		Symbols: symbols,
		Type:    requestType,
		From:    from,
	}
}

func (m *MsgRequestPrice) ValidateBasic() error {
	if len(m.Symbols) == 0 {
		return fmt.Errorf("no symbols requested")
	}

	return nil
}

func (m *MsgRequestPrice) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
