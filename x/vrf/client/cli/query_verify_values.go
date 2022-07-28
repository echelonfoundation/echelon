package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/echelonfoundation/echelon/v3/x/vrf/types"
)

var _ = strconv.Itoa(0)

func CmdVerifyValues() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-values [pubkey] [message] [vrv] [proof]",
		Short: "Query verify-values",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPubkey := args[0]
			reqMessage := args[1]
			reqVrv := args[2]
			reqProof := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryVerifyValuesRequest{

				Pubkey:  reqPubkey,
				Message: reqMessage,
				Vrv:     reqVrv,
				Proof:   reqProof,
			}

			res, err := queryClient.VerifyValues(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
