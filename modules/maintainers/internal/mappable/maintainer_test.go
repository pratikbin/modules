/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	qualifiedMappables "github.com/persistenceOne/persistenceSDK/schema/mappables/qualified"
	"testing"

	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	qualifiedTraits "github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
)

func Test_Maintainer_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID"), base.NewStringData("MutableData")))

	testMaintainerID := key.NewMaintainerID(classificationID, identityID)

	testMaintainer := NewMaintainer(testMaintainerID, nil, mutableProperties).(maintainer)

	require.Equal(t, maintainer{Document: qualifiedMappables.Document{ID: testMaintainerID, ClassificationID: classificationID, HasImmutables: qualifiedTraits.HasImmutables{Properties: immutableProperties}, HasMutables: qualifiedTraits.HasMutables{Properties: mutableProperties}}}, testMaintainer)
	require.Equal(t, testMaintainerID, testMaintainer.GetID())
	require.Equal(t, classificationID, testMaintainer.GetClassificationID())
	require.Equal(t, identityID, testMaintainer.GetIdentityID())
	require.Equal(t, true, testMaintainer.CanAddMaintainer())
	require.Equal(t, true, testMaintainer.CanMutateMaintainer())
	require.Equal(t, true, testMaintainer.CanRemoveMaintainer())
	require.Equal(t, true, testMaintainer.MaintainsProperty(base.NewID("ID")))
	require.Equal(t, false, testMaintainer.MaintainsProperty(base.NewID("ID2")))
	require.Equal(t, testMaintainerID, testMaintainer.GetKey())
}
