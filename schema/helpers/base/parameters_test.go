// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	baseTestUtilities "github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
)

func TestParameters(t *testing.T) {
	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	Parameter := baseTypes.NewParameter(baseIDs.NewID("testParameter"), baseData.NewStringData("testData"), func(interface{}) error { return nil })
	ParameterList := []types.Parameter{Parameter}
	Parameters := NewParameters(ParameterList...)
	subspace := params.NewSubspace(codec, storeKey, transientStoreKey, "test").WithKeyTable(Parameters.GetKeyTable())
	subspace.SetParamSet(context, Parameters)
	Parameters = Parameters.Initialize(subspace).(parameters)

	require.NotNil(t, Parameters.ParamSetPairs())

	require.NotNil(t, Parameters.GetKeyTable())

	require.Equal(t, true, Parameters.Equal(Parameters))

	require.Equal(t, true, Parameters.GetList()[0].Equal(Parameter))
	require.Equal(t, `{"id":{"idString":"testParameter"},"data":{"value":"testData"}}`, Parameters.String())

	err := Parameters.Validate()
	require.Nil(t, err)

	require.NotPanics(t, func() {
		Parameters.Fetch(context, baseIDs.NewID("testParameter"))
	})

	require.Equal(t, "testData123", Parameters.Mutate(context,
		baseTypes.NewParameter(baseIDs.NewID("testParameter"), baseData.NewStringData("testData123"), func(interface{}) error { return nil })).Get(baseIDs.NewID("testParameter")).GetData().String())
}
