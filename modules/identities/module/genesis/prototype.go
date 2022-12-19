// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package genesis

import (
	"github.com/AssetMantle/modules/modules/identities/module/key"
	"github.com/AssetMantle/modules/modules/identities/module/mappable"
	"github.com/AssetMantle/modules/modules/identities/module/parameters"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

// TODO ***** add module nub IDs
func Prototype() helpers.Genesis {
	return baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, []helpers.Mappable{}, parameters.Prototype().GetList())
}
