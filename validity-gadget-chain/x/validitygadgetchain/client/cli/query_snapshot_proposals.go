package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

var _ = strconv.Itoa(0)

func CmdSnapshotProposals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot-proposals",
		Short: "Query snapshot_proposals",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySnapshotProposalsRequest{}

			res, err := queryClient.SnapshotProposals(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
