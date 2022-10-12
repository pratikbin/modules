// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/traits"
)

type metaProperty struct {
	ID   ids.PropertyID `json:"id"`
	Data data.Data      `json:"data"`
}

var _ properties.MetaProperty = (*metaProperty)(nil)

func (metaProperty metaProperty) GetData() data.Data {
	return metaProperty.Data
}
func (metaProperty metaProperty) RemoveData() properties.Property {
	return NewMesaProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty metaProperty) GetID() ids.PropertyID {
	return metaProperty.ID
}
func (metaProperty metaProperty) GetDataID() ids.DataID {
	return metaProperty.Data.GetID()
}
func (metaProperty metaProperty) GetKey() ids.StringID {
	return metaProperty.ID.GetKey()
}
func (metaProperty metaProperty) GetType() ids.StringID {
	return metaProperty.Data.GetType()
}
func (metaProperty metaProperty) IsMeta() bool {
	return true
}
func (metaProperty metaProperty) Compare(listable traits.Listable) int {
	if compareMetaProperty, err := metaPropertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return metaProperty.GetID().Compare(compareMetaProperty.GetID())
	}
}
func metaPropertyFromInterface(listable traits.Listable) (metaProperty, error) {
	switch value := listable.(type) {
	case metaProperty:
		return value, nil
	default:
		return metaProperty{}, errorConstants.MetaDataError
	}
}

func NewEmptyMetaPropertyFromID(propertyID ids.PropertyID) properties.MetaProperty {
	return metaProperty{
		ID: propertyID,
	}
}
func NewMetaProperty(key ids.StringID, data data.Data) properties.MetaProperty {
	if data == nil || key == nil {
		panic(errorConstants.MetaDataError)
	}
	return metaProperty{
		ID:   baseIDs.NewPropertyID(key, data.GetType()),
		Data: data,
	}
}
