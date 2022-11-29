// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"github.com/AssetMantle/modules/modules/metas/auxiliaries"
	"github.com/AssetMantle/modules/modules/metas/internal/block"
	"github.com/AssetMantle/modules/modules/metas/internal/genesis"
	"github.com/AssetMantle/modules/modules/metas/internal/mapper"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/modules/metas/internal/parameters"
	"github.com/AssetMantle/modules/modules/metas/internal/queries"
	"github.com/AssetMantle/modules/modules/metas/internal/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		1,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		nil,
		mapper.Prototype,
		nil,
		parameters.Prototype,
		queries.Prototype,
		nil,
		transactions.Prototype,
	)
}
