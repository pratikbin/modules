/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

var Prototype = base.NewGenesis(packageCodec, []traits.Mappable{}, parameters.Prototype.GetList())
