// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/types/base"
	"github.com/AssetMantle/modules/utilities/meta"
)

func Test_StringData(t *testing.T) {
	value := "data"
	testStringData := NewStringData(value)
	testStringData2 := NewStringData("")

	require.Equal(t, value, testStringData.String())
	require.Equal(t, base.NewID(meta.Hash(value)), testStringData.GenerateHashID())
	require.Equal(t, base.NewID(""), testStringData2.GenerateHashID())
	require.Equal(t, StringDataID, testStringData.GetTypeID())
	require.Equal(t, testStringData.ZeroValue(), NewStringData(""))

	data, err := ReadStringData("testString")
	require.Nil(t, err)
	require.Equal(t, stringData{Value: "testString"}.String(), data.String())

	require.Equal(t, false, testStringData.Compare(testStringData2) == 0)
	require.Equal(t, true, testStringData.Compare(testStringData) == 0)
	require.Panics(t, func() {
		require.Equal(t, false, testStringData.Compare(NewIDData(base.NewID("ID"))) == 0)
	})
}
