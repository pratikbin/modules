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

func (assetID *AssetID) Compare(listable traits.Listable) int {
	return bytes.Compare(assetID.Bytes(), assetIDFromInterface(listable).Bytes())

}
func (assetID *AssetID) IsOwnableID() {}
func (assetID *AssetID) IsAssetID()   {}
func (assetID *AssetID) Bytes() []byte {
	return assetID.AssetID.HashId.GetHashID().IdBytes
}
func assetIDFromInterface(i interface{}) *AssetIDI {
	switch value := i.(type) {
	case *AssetIDI:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func GenerateAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return NewAssetID(GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()))
}

func NewAssetID(hashID ids.HashID) ids.AssetID {
	return &AssetID{
		HashId: hashID,
	}
}

func PrototypeAssetID() ids.AssetID {
	return NewAssetID(PrototypeHashID())
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if hashID, err := ReadHashID(assetIDString); err == nil {
		return NewAssetID(hashID), nil
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return PrototypeAssetID(), errorConstants.MetaDataError
}
