/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

type immutables struct {
	Properties types.Properties `json:"properties"`
}

var _ types.Immutables = (*immutables)(nil)

func (immutables immutables) Get() types.Properties {
	return immutables.Properties
}
func (immutables immutables) GetHashID() types.ID {
	if len(immutables.Properties.GetList()) == 0 {
		return NewID("")
	}

	metaList := make([]string, len(immutables.Properties.GetList()))
	for i, immutableProperty := range immutables.Properties.GetList() {
		metaList[i] = immutableProperty.GetFact().GetHashID().String()
	}

	return NewID(metaUtilities.Hash(metaList...))
}
func NewImmutables(properties types.Properties) types.Immutables {
	return immutables{Properties: properties}
}
