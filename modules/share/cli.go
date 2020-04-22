package share

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/persistenceOne/persistenceSDK/modules/share/queries/share"
	"github.com/persistenceOne/persistenceSDK/modules/share/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/share/transactions/lock"
	"github.com/persistenceOne/persistenceSDK/modules/share/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/share/transactions/send"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCLIRootTransactionCommand(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        TransactionRoute,
		Short:                      "Share root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootTransactionCommand.AddCommand(flags.PostCommands(
		burn.TransactionCommand(codec),
		lock.TransactionCommand(codec),
		mint.TransactionCommand(codec),
		send.TransactionCommand(codec),
	)...)
	return rootTransactionCommand
}

func GetCLIRootQueryCommand(codec *codec.Codec) *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        QuerierRoute,
		Short:                      "Share root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootQueryCommand.AddCommand(flags.GetCommands(
		share.QueryCommand(codec),
	)...)
	return rootQueryCommand
}
