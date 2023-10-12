package cli

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ojo-network/ojo-chainibc-demo/x/oracleintegration/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
	flagRequestType            = "request-type"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetMsgRequestPriceCmd())

	return cmd
}

func GetMsgRequestPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-price [denoms,]",
		Args:  cobra.MinimumNArgs(2),
		Short: "request price for denoms",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestType, err := cmd.Flags().GetInt32(flagRequestType)
			if err != nil {
				return err
			}

			msg := types.NewMsgPriceRequest(
				args,
				types.PriceRequestType(requestType),
				clientCtx.GetFromAddress().String(),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Int32(flagRequestType, 0, "request type for prices")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
