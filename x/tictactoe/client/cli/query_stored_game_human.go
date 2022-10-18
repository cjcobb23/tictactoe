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

func CmdStoredGameHuman() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stored-game-human [game-index]",
		Short: "Query storedGameHuman",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqGameIndex, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStoredGameHumanRequest{

				GameIndex: reqGameIndex,
			}

			res, err := queryClient.StoredGameHuman(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
