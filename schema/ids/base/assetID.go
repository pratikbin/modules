// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

var _ ids.AssetID = (*AssetID)(nil)

func (assetID AssetID) Compare(listable traits.Listable) int {
	return bytes.Compare(assetID.Bytes(), *assetIDFromInterface(listable).HashId)

}

func (assetID AssetID) IsOwnableID() {}
func (assetID AssetID) IsAssetID()   {}

func assetIDFromInterface(i interface{}) *AssetID {
	switch value := i.(type) {
	case AssetID:
		return &value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func (assetID AssetID) Bytes() []byte {
	return *assetID.HashId
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return &AssetID{
		HashId: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeAssetID() ids.AssetID {
	return &AssetID{
		HashId: PrototypeHashID().(*HashID),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return &AssetID{
			HashId: hashID.(*HashID),
		}, nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return &AssetID{}, errorConstants.MetaDataError
}
