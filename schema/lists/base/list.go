// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/traits/base"
	"sort"

	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO type check that list is same type and test cases
//type list []traits.Listable

var _ lists.List = (*List)(nil)

func (list List) Get() []traits.Listable {
	var listables []traits.Listable
	for _, i := range list.Listables {
		listables = append(listables, i)
	}
	return listables
}

//func (list List) Size() int {
//	return len(list)
//}
func (list List) Search(listable traits.Listable) (int, bool) {
	index := sort.Search(
		len(list.Listables),
		func(i int) bool {
			return list.Listables[i].Compare(listable) >= 0
		},
	)

	if index < len(list.Listables) && list.Listables[index].Compare(listable) == 0 {
		return index, true
	}

	return index, false
}
func (list List) Add(listables ...traits.Listable) lists.List {
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); !found {
			updatedList.Listables = append(updatedList.Listables, listable.(*base.Listable))
			copy(updatedList.Listables[index+1:], updatedList.Listables[index:])
			updatedList.Listables[index] = listable.(*base.Listable)
		}
	}

	return &updatedList
}
func (list List) Remove(listables ...traits.Listable) lists.List {
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); found {
			updatedList.Listables = append(updatedList.Listables[:index], updatedList.Listables[index+1:]...)
		}
	}

	return &updatedList
}
func (list List) Mutate(listables ...traits.Listable) lists.List {
	// TODO write test
	updatedList := list

	for _, listable := range listables {
		if index, found := updatedList.Search(listable); found {
			updatedList.Listables[index] = listable.(*base.Listable)
		}
	}

	return &updatedList
}

func NewList(listables ...traits.Listable) lists.List {
	list := List(listables)
	sort.Slice(list, func(i, j int) bool {
		return list[i].Compare(list[j]) <= 0
	})

	return &list
}
