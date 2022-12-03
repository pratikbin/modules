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
	ids.PropertyID
	data.DataI
}

var _ properties.MetaProperty = (*metaProperty)(nil)

func (metaProperty metaProperty) GetData() data.DataI {
	return metaProperty.DataI
}
func (metaProperty metaProperty) ScrubData() properties.MesaProperty {
	return NewMesaProperty(metaProperty.GetKey(), metaProperty.GetData())
}
func (metaProperty metaProperty) GetID() ids.PropertyID {
	return metaProperty.PropertyID
}
func (metaProperty metaProperty) GetDataID() ids.DataID {
	return metaProperty.DataI.GetID()
}
func (metaProperty metaProperty) GetKey() ids.StringID {
	return metaProperty.PropertyID.GetKey()
}
func (metaProperty metaProperty) GetType() ids.StringID {
	return metaProperty.DataI.GetType()
}
func (metaProperty metaProperty) IsMeta() bool {
	return true
}
func (metaProperty metaProperty) Compare(listable traits.Listable) int {
	// NOTE: compare property can be meta or mesa, so listable must only be cast to Property Interface
	if compareProperty, err := propertyFromInterface(listable); err != nil {
		panic(err)
	} else {
		return metaProperty.GetID().Compare(compareProperty.GetID())
	}
}

func NewEmptyMetaPropertyFromID(propertyID ids.PropertyID) properties.MetaProperty {
	return metaProperty{
		PropertyID: propertyID,
	}
}
func NewMetaProperty(key ids.StringID, data data.DataI) properties.MetaProperty {
	if data == nil || key == nil {
		panic(errorConstants.MetaDataError)
	}
	return metaProperty{
		PropertyID: baseIDs.NewPropertyID(key, data.GetType()),
		DataI:      data,
	}
}
