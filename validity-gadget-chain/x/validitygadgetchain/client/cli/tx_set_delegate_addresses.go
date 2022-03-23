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

func CmdSetDelegateAddresses() *cobra.Command {
	cmd := &cobra.Command{
		Use: "set-delegate-addresses [validator] [btc_oracle]	[eth_oracle] [watch_tower]",
		Short: "Broadcast message setDelegateAddresses",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argValidator := args[0]
			argBTCOracle := args[1]
			argETHOracle := args[2]
			argWatchtower := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetDelegateAddresses(
				clientCtx.GetFromAddress().String(),
				argValidator,
				argBTCOracle,
				argETHOracle,
				argWatchtower,
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
