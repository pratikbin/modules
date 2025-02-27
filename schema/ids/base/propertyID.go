// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

//type propertyID struct {
//	Key  ids.StringID
//	Type ids.StringID
//}

var _ ids.PropertyID = (*PropertyID)(nil)

func (propertyID *PropertyID) IsPropertyID() {}
func (propertyID *PropertyID) GetKey() ids.StringID {
	return propertyID.KeyID
}
func (propertyID *PropertyID) GetType() ids.StringID {
	return propertyID.TypeID
}
func (propertyID *PropertyID) PropertyIDString() string {
	return stringUtilities.JoinIDStrings(propertyID.KeyID.String(), propertyID.TypeID.String())
}
func (propertyID *PropertyID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, propertyID.KeyID.Bytes()...)
	Bytes = append(Bytes, propertyID.TypeID.Bytes()...)

	return Bytes
}
func (propertyID *PropertyID) Compare(listable traits.Listable) int {
	return bytes.Compare(propertyID.Bytes(), propertyIDFromInterface(listable).Bytes())
}
func (propertyID *PropertyID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_PropertyID{
			PropertyID: propertyID,
		},
	}
}

func propertyIDFromInterface(listable traits.Listable) *PropertyID {
	switch value := listable.(type) {
	case *PropertyID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}

func NewPropertyID(key, Type ids.StringID) ids.PropertyID {
	return &PropertyID{
		KeyID:  key.(*StringID),
		TypeID: Type.(*StringID),
	}
}
