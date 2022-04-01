// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package traits

import "github.com/persistenceOne/persistenceSDK/schema/types"

type Ownable interface {
	GetOwnerID() types.ID
	GetOwnableID() types.ID
}
