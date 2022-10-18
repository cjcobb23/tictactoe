package cli

import (
	"strconv"

	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "move [game-index] [x] [y]",
		Short: "Broadcast message move",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argGameIndex, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argX, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			argY, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMove(
				clientCtx.GetFromAddress().String(),
				argGameIndex,
				argX,
				argY,
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
