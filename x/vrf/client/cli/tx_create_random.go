package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

var _ = strconv.Itoa(0)

func CmdCreateRandom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-random [multiplier]",
		Short: "create-random",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argMultiplier, erri := strconv.ParseUint(args[0], 10, 64)
			if erri != nil {
				return erri
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateRandom(
				clientCtx.GetFromAddress().String(),
				argMultiplier,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
