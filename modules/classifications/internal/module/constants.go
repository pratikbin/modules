// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

const Name = "classifications"
const ConsensusVersion = 1

var StoreKeyPrefix = constants.ClassificationsStoreKeyPrefix

// MaxPropertyCount TODO convert it to module param
const MaxPropertyCount = 22
