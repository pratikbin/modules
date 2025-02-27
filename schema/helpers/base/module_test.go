// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"encoding/json"
	"math/rand"
	"testing"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModule "github.com/cosmos/cosmos-sdk/types/module"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	parameters2 "github.com/AssetMantle/modules/schema/parameters"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
	"github.com/AssetMantle/modules/utilities/test"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

var auxiliariesPrototype = func() helpers.Auxiliaries {
	return auxiliaries{[]helpers.Auxiliary{NewAuxiliary("testAuxiliary", baseTestUtilities.TestAuxiliaryKeeperPrototype)}}
}
var genesisPrototype = func() helpers.Genesis {
	return NewGenesis(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype,
		[]helpers.Mappable{baseTestUtilities.NewMappable("test", "testValue")},
		[]parameters2.Parameter{baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })})
}
var mapperPrototype = func() helpers.Mapper {
	return NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype)
}
var parametersPrototype = func() helpers.Parameters {
	return NewParameters(baseTypes.NewParameter(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil }))
}
var queriesPrototype = func() helpers.Queries {
	return queries{[]helpers.Query{NewQuery("testQuery", "q", "testQuery", "test", baseTestUtilities.TestQueryRequestPrototype,
		baseTestUtilities.TestQueryResponsePrototype, baseTestUtilities.TestQueryKeeperPrototype)}}
}
var simulatorPrototype = func() helpers.Simulator { return nil }
var transactionsPrototype = func() helpers.Transactions {
	return transactions{[]helpers.Transaction{NewTransaction("TestMessage", "", "", baseTestUtilities.TestTransactionRequestPrototype, baseTestUtilities.TestMessagePrototype,
		baseTestUtilities.TestTransactionKeeperPrototype)}}
}
var blockPrototype = func() helpers.Block { return baseTestUtilities.TestBlockPrototype() }

func TestModule(t *testing.T) {
	context, storeKey, transientStoreKey := test.SetupTest(t)
	var legacyAmino = sdkCodec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	subspace := paramsTypes.NewSubspace(legacyAmino, storeKey, transientStoreKey, "test") // .WithKeyTable(parametersPrototype().GetKeyTable())
	// subspace.SetParamSet(context, parametersPrototype())
	Module := NewModule("test", auxiliariesPrototype, genesisPrototype,
		mapperPrototype, parametersPrototype, queriesPrototype, simulatorPrototype, transactionsPrototype, blockPrototype).Initialize(storeKey, subspace).(module)

	// AppModuleBasic
	require.Equal(t, "test", Module.Name())

	// RegisterLegacyAminoCodec
	Module.RegisterLegacyAminoCodec(legacyAmino)

	require.NotPanics(t, func() {
		Module.DefaultGenesis()
	})

	require.NotPanics(t, func() {
	})
	require.Nil(t, Module.ValidateGenesis(Module.DefaultGenesis()))

	router := mux.NewRouter()
	require.NotPanics(t, func() {
		Module.RegisterRESTRoutes(context, router)
	})

	// GetTxCmd
	require.Equal(t, "test", Module.GetTxCmd(legacyAmino).Name())
	require.Equal(t, "test", Module.GetQueryCmd(legacyAmino).Name())

	// AppModule
	require.NotPanics(t, func() {
		Module.RegisterInvariants(nil)
	})
	require.Equal(t, "test", Module.Route())

	response, err := Module.NewHandler()(context, baseTestUtilities.NewTestMessage(sdkTypes.AccAddress("addr"), "id"))
	require.Nil(t, err)
	require.NotNil(t, response)

	require.Equal(t, "test", Module.QuerierRoute())

	encodedRequest, err := Module.queries.Get("testQuery").(query).requestPrototype().Encode()
	require.Nil(t, err)

	queryResponse, err := Module.NewQuerierHandler()(context, []string{"testQuery"}, abciTypes.RequestQuery{Data: encodedRequest})
	require.Nil(t, err)
	require.NotNil(t, queryResponse)

	require.NotPanics(t, func() {
		Module.BeginBlock(context, abciTypes.RequestBeginBlock{})
	})
	endBlockResponse := Module.EndBlock(context, abciTypes.RequestEndBlock{})
	require.Equal(t, []abciTypes.ValidatorUpdate{}, endBlockResponse)

	require.NotPanics(t, func() {
		Module.InitGenesis(context, Module.DefaultGenesis())
	})

	require.Equal(t, Module.DefaultGenesis(), Module.ExportGenesis(context))
	// AppModuleSimulation
	require.Panics(t, func() {
		Module.GenerateGenesisState(&sdkModule.SimulationState{})
		Module.ProposalContents(sdkModule.SimulationState{})
		Module.RandomizedParams(&rand.Rand{})
		Module.RegisterStoreDecoder(sdkTypes.StoreDecoderRegistry{})
		Module.WeightedOperations(sdkModule.SimulationState{})
	})

	// types.Module
	require.Equal(t, "testAuxiliary", Module.GetAuxiliary("testAuxiliary").GetName())
	_, err = Module.DecodeModuleTransactionRequest("TestMessage", json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Nil(t, err)
}
