package base

import "github.com/AssetMantle/modules/schema/traits"

var _ (traits.Listable) = (*Listable)(nil)

func (m *Listable) Compare(listable traits.Listable) int {
	return m.Listable.Compare(listable)
}
