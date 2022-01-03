/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
)

func Test_MetaProperty(t *testing.T) {
	metaFact1 := NewMetaFact(NewHeightData(NewHeight(123)))
	testMetaProperty := NewMetaProperty(NewID("ID"), metaFact1)
	require.Equal(t, metaFact1, testMetaProperty.GetMetaFact())
	require.Equal(t, NewProperty(NewID("ID"), NewFact(NewHeightData(NewHeight(123)))), testMetaProperty.RemoveData())
	require.Equal(t, NewID("ID"), testMetaProperty.GetID())
	require.Equal(t, metaFact1, testMetaProperty.GetMetaFact())

	readMetaProperty, err := ReadMetaProperty("ID2:S|SomeData")
	require.Equal(t, NewMetaProperty(NewID("ID2"), NewMetaFact(NewStringData("SomeData"))), readMetaProperty)
	require.Nil(t, err)

	readMetaProperty, err = ReadMetaProperty("RandomValue")
	require.Equal(t, nil, readMetaProperty)
	require.Equal(t, errors.IncorrectFormat, err)

	readMetaProperty, err = ReadMetaProperty("RandomID:RandomValue")
	require.Equal(t, nil, readMetaProperty)
	require.Equal(t, errors.IncorrectFormat, err)
}
