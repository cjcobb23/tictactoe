package cli

import (
	"strconv"

	"github.com/cjcobb23/tictactoe/x/tictactoe/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCanPlayMove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "can-play-move [game-index] [player] [x] [y]",
		Short: "Query canPlayMove",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGameIndex, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			reqPlayer := args[1]
			reqX, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			reqY, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCanPlayMoveRequest{

				GameIndex: reqGameIndex,
				Player:    reqPlayer,
				X:         reqX,
				Y:         reqY,
			}

			res, err := queryClient.CanPlayMove(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
