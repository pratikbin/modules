package base

import "C"
import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type classificationID struct {
//	ids.HashID
// }

var _ ids.ClassificationID = (*ClassificationID)(nil)

func (classificationID ClassificationID) Bytes() []byte {
	return classificationID.HashId.Bytes()
}

func (classificationID ClassificationID) IsClassificationID() {}

func (classificationID ClassificationID) Compare(listable traits.Listable) int {
	return bytes.Compare(classificationID.Bytes(), classificationIDFromInterface(listable).Bytes())
}

func classificationIDFromInterface(i interface{}) ClassificationID {
	switch value := i.(type) {
	case ClassificationID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}

func NewClassificationID(immutables qualified.Immutables, mutables qualified.Mutables) ids.ClassificationID {
	immutableIDByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))

	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		immutableIDByteList[i] = property.GetID().Bytes()
	}

	mutableIDByteList := make([][]byte, len(mutables.GetMutablePropertyList().GetList()))

	for i, property := range mutables.GetMutablePropertyList().GetList() {
		mutableIDByteList[i] = property.GetID().Bytes()
	}

	defaultImmutableByteList := make([][]byte, len(immutables.GetImmutablePropertyList().GetList()))

	for i, property := range immutables.GetImmutablePropertyList().GetList() {
		// TODO test
		if hashID := property.GetDataID().GetHashID(); !(hashID.Compare(GenerateHashID()) == 0) {
			defaultImmutableByteList[i] = hashID.Bytes()
		}
	}

	return &ClassificationID{HashId: GenerateHashID(GenerateHashID(immutableIDByteList...).Bytes(), GenerateHashID(mutableIDByteList...).Bytes(), GenerateHashID(defaultImmutableByteList...).Bytes()).(*HashID)}
}

func PrototypeClassificationID() ids.ClassificationID {
	return &ClassificationID{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadClassificationID(classificationIDString string) (ids.ClassificationID, error) {
	if hashID, err := ReadHashID(classificationIDString); err == nil {
		return &ClassificationID{HashId: hashID.(*HashID)}, nil
	}

	if classificationIDString == "" {
		return PrototypeClassificationID(), nil
	}

	return &ClassificationID{}, constants.MetaDataError
}
