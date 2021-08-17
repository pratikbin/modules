/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Query interface {
	GetName() string
	Command() *cobra.Command
	HandleMessageByLegacyAmino(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(client.Context) http.HandlerFunc
	Initialize(Mapper, Parameters, ...interface{}) Query
}
