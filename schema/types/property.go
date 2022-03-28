/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

// TODO do sortable interface
type Property interface {
	GetID() ID
	GetDataID() ID
	GetKeyID() ID
	GetTypeID() ID
	GetHashID() ID
}
