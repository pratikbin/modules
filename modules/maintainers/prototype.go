// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries"
	"github.com/AssetMantle/modules/modules/maintainers/internal/block"
	"github.com/AssetMantle/modules/modules/maintainers/internal/genesis"
	"github.com/AssetMantle/modules/modules/maintainers/internal/invariants"
	"github.com/AssetMantle/modules/modules/maintainers/internal/mapper"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	"github.com/AssetMantle/modules/modules/maintainers/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/internal/queries"
	"github.com/AssetMantle/modules/modules/maintainers/internal/simulator"
	"github.com/AssetMantle/modules/modules/maintainers/internal/transactions"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Module {
	return baseHelpers.NewModule(
		module.Name,
		module.ConsensusVersion,
		auxiliaries.Prototype,
		block.Prototype,
		genesis.Prototype,
		invariants.Prototype,
		mapper.Prototype,
		parameters.Prototype,
		queries.Prototype,
		simulator.Prototype,
		transactions.Prototype,
	)
}
