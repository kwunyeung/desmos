package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/kwunyeung/desmos/x/magpie/types"
)

// GetTxCmd set the tx commands
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	magpieTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "magpie transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	magpieTxCmd.AddCommand(client.PostCommands(
		GetCmdCreateSession(cdc),
		client.LineBreak,
		GetCmdCreatePost(cdc),
		GetCmdEditPost(cdc),
		GetCmdAddLike(cdc),
	)...)

	return magpieTxCmd
}

// GetCmdCreatePost is the CLI command for creating a post
func GetCmdCreatePost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create [message] [parent-post-id] [namespace] [external address]",
		Short: "create a new post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)

			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			msg := types.NewMsgCreatePost(args[0], args[1], time.Now(), from, args[2], args[3])
			var err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdEditPost is the CLI command for editing a post
func GetCmdEditPost(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "edit [post-id] [message]",
		Short: "edit an owned post",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)

			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			msg := types.NewMsgEditPost(args[0], args[1], time.Now(), from)
			var err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdAddLike is the CLI command for adding a like to a post
func GetCmdAddLike(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "like [post-id] [namespace] [external address]",
		Short: "like a post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			accGetter := authtypes.NewAccountRetriever(cliCtx)

			from := cliCtx.GetFromAddress()
			if err := accGetter.EnsureExists(from); err != nil {
				return err
			}

			msg := types.NewMsgLike(args[0], time.Now(), from, args[1], args[2])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdCreateSession is the CLI command for creating a session for create post
func GetCmdCreateSession(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-session [namespace] [external address] [pubkey] [external signer signature]",
		Short: "record a session for external service to post a magpie",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			// accGetter := authtypes.NewAccountRetriever(cliCtx)

			from := cliCtx.GetFromAddress()
			// if err := accGetter.EnsureExists(from); err != nil {
			// 	return err
			// }

			// pubkey, _ := sdk.GetAccPubKeyBech32(args[2])

			msg := types.NewMsgCreateSession(time.Now(), from, args[0], args[1], args[2], args[3])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			// cliCtx.PrintResponse = true

			// return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
