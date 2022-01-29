package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/twilight-project/validity-gadget-chain/x/validitygadgetchain/types"
)

var _ = strconv.Itoa(0)

func CmdSnapshotProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot-proposal [iteration] [log-snap]",
		Short: "Broadcast message SnapshotProposal",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argIteration := args[0]
			argLogSnap := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSnapshotProposal(
				clientCtx.GetFromAddress().String(),
				argIteration,
				argLogSnap,
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
