package mappable

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Meta_Methods(t *testing.T) {
	data := base.NewStringData("Data")

	testMeta := NewMeta(data).(meta)
	require.Equal(t, meta{ID: key.NewMetaID(data.GetTypeID(), data.GenerateHashID()), Data: data}, testMeta)
	require.Equal(t, data, testMeta.GetData())
	require.Equal(t, key.NewMetaID(data.GetTypeID(), data.GenerateHashID()), testMeta.GetKey())
	require.Equal(t, key.GenerateMetaID(data), testMeta.GetID())

}
