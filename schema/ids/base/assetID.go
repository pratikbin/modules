// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type assetID struct {
	ids.ClassificationID
	ids.HashID
}

var _ ids.AssetID = (*assetID)(nil)

func (assetID assetID) String() string {
	return stringUtilities.JoinIDStrings(assetID.ClassificationID.String(), assetID.HashID.String())
}
func (assetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.HashID.Bytes()...)

	return Bytes
}
func (assetID assetID) Compare(listable traits.Listable) int {
	if compareAssetID, err := assetIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return bytes.Compare(assetID.Bytes(), compareAssetID.Bytes())
	}
}
func assetIDFromInterface(i interface{}) (assetID, error) {
	switch value := i.(type) {
	case assetID:
		return value, nil
	default:
		return assetID{}, errorConstants.MetaDataError
	}
}
func NewAssetID(classificationID ids.ClassificationID, immutableProperties lists.PropertyList) ids.AssetID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           base.Immutables{PropertyList: immutableProperties}.GenerateHashID(),
	}
}
