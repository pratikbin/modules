package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type metaID struct {
	Type ids.StringID
	ids.HashID
}

var _ ids.MetaID = (*metaID)(nil)

func (metaID metaID) IsMetaID() {}
func (metaID metaID) String() string {
	return stringUtilities.JoinIDStrings(metaID.Type.String(), metaID.HashID.String())
}
func (metaID metaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, metaID.Type.Bytes()...)
	Bytes = append(Bytes, metaID.HashID.Bytes()...)

	return Bytes
}

// TODO compare components
func (metaID metaID) Compare(listable traits.Listable) int {
	return bytes.Compare(metaID.Bytes(), metaIDFromInterface(listable).Bytes())
}

// TODO return error and not panic for all
func metaIDFromInterface(i interface{}) metaID {
	switch value := i.(type) {
	case metaID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewMetaID(Type ids.StringID, hashID ids.HashID) ids.MetaID {
	return metaID{
		Type:   Type,
		HashID: hashID,
	}
}

func PrototypeMetaID() ids.MetaID {
	return metaID{
		Type:   PrototypeStringID(),
		HashID: PrototypeHashID(),
	}
}

func ReadMetaID(metaIDString string) (ids.MetaID, error) {
	if metaIDString := stringUtilities.SplitCompositeIDString(metaIDString); len(metaIDString) == 2 {
		Type := NewStringID(metaIDString[0])
		if hashID, err := ReadHashID(metaIDString[1]); err == nil {
			return metaID{
				Type:   Type,
				HashID: hashID,
			}, nil
		}
	}

	if metaIDString == "" {
		return PrototypeMetaID(), nil
	}

	return metaID{}, constants.MetaDataError
}
