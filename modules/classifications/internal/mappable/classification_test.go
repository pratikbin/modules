/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"

	baseTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/classifications/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Classification_Methods(t *testing.T) {

	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	chainID := base.NewID("chainID")
	id := key.NewClassificationID(chainID, immutableProperties, mutableProperties)

	testClassification := NewClassification(id, immutableProperties, mutableProperties)
	require.Equal(t, classification{Document: qualified.Document{ID: id, HasImmutables: baseTraits.HasImmutables{Properties: immutableProperties}, HasMutables: baseTraits.HasMutables{Properties: mutableProperties}}}, testClassification)
	require.Equal(t, immutableProperties, testClassification.GetImmutableProperties())
	require.Equal(t, mutableProperties, testClassification.GetMutableProperties())
	require.Equal(t, key.FromID(id), testClassification.GetKey())
	require.Equal(t, id, testClassification.(classification).GetID())
}
