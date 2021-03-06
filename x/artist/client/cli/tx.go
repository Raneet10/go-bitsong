package cli

import (
	"bufio"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/artist/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"strings"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	contentTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	contentTxCmd.AddCommand(flags.PostCommands(
		GetCmdArtistCreate(cdc),
	)...)

	return contentTxCmd
}

func GetCmdArtistCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new artist on bitsong",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new artist on bitsong.
Example:
$ %s tx artist create [id] [name] --from <owner>`,
				version.ClientName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			id := args[0]
			name := args[1]

			msg := types.NewMsgArtistCreate(id, name, nil, nil, "", cliCtx.FromAddress)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
