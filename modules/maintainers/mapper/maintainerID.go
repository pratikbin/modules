/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"bytes"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type maintainerID struct {
	MaintainedID types.ID `json:"maintainedID" valid:"required~required field maintainedID missing"`
	IdentityID   types.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ types.ID = (*maintainerID)(nil)

func (maintainerID maintainerID) Bytes() []byte {
	return append(
		maintainerID.MaintainedID.Bytes(),
		maintainerID.IdentityID.Bytes()...)

}

func (maintainerID maintainerID) String() string {
	var values []string
	values = append(values, maintainerID.MaintainedID.String())
	values = append(values, maintainerID.IdentityID.String())
	return strings.Join(values, constants.CompositeIDSeparator)
}

func (maintainerID maintainerID) Compare(id types.ID) int {
	return bytes.Compare(maintainerID.Bytes(), id.Bytes())
}

func readMaintainerID(maintainerIDString string) types.ID {
	idList := strings.Split(maintainerIDString, constants.CompositeIDSeparator)
	if len(idList) == 2 {
		return maintainerID{
			MaintainedID: base.NewID(idList[0]),
			IdentityID:   base.NewID(idList[1]),
		}
	}
	return maintainerID{IdentityID: base.NewID(""), MaintainedID: base.NewID("")}
}

func maintainerIDFromInterface(id types.ID) maintainerID {
	switch value := id.(type) {
	case maintainerID:
		return value
	default:
		return maintainerIDFromInterface(readMaintainerID(id.String()))
	}
}
func generateKey(maintainerID types.ID) []byte {
	return append(StoreKeyPrefix, maintainerIDFromInterface(maintainerID).Bytes()...)
}
func NewMaintainerID(maintainedID types.ID, identityID types.ID) types.ID {
	return maintainerID{
		MaintainedID: maintainedID,
		IdentityID:   identityID,
	}
}
