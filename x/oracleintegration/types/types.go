package types

import (
	"strings"
)

func (msg *MsgRequestPrice) GetRequestedSymbols() []string {
	symbols := msg.GetSymbols()
	for i, symbol := range symbols {
		symbols[i] = strings.ToUpper(symbol)
	}

	return symbols
}
