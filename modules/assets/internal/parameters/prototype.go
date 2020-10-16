/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package parameters

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

func Prototype() helpers.Parameters {
	return base.NewParameters(dummy.Parameter)
}
