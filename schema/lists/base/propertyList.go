package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

//type propertyList struct {
//	lists.List
//}

var _ lists.PropertyList = (*PropertyList)(nil)

func (propertyList *PropertyList) GetProperty(propertyID ids.PropertyID) properties.Property {
	list := propertyList.GetPropertyList()
	if i, found := list.(lists.List).Search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}

	return nil
}

func (propertyList *PropertyList) GetPropertyList() []properties.Property {
	Properties := make([]properties.Property, len(propertyList.GetList()))
	for i, listable := range propertyList.GetList() {
		Properties[i] = listable.Impl.(properties.Property)
	}
	return Properties
}

func (propertyList *PropertyList) GetPropertyIDList() lists.IDList {
	propertyIDList := NewIDList()
	for _, property := range propertyList.GetList() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}
	return propertyIDList
}
func (propertyList *PropertyList) Add(properties ...properties.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Add(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList *PropertyList) Remove(properties ...properties.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Remove(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList *PropertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	propertyList.List = propertyList.List.Mutate(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList *PropertyList) ScrubData() lists.PropertyList {
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.List.Get() {
		if property := listable.(properties.Property); property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.(properties.MetaProperty).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}
func propertiesToListables(properties ...properties.Property) []traits.Listable {
	listables := make([]traits.Listable, len(properties))
	for i, property := range properties {
		listables[i] = property
	}
	return listables
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	return propertyList{List: NewList(propertiesToListables(properties...)...)}
}
