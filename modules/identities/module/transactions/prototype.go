// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"github.com/AssetMantle/modules/modules/identities/module/transactions/define"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/deputize"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/issue"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/mutate"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/nub"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/provision"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/quash"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/revoke"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/unprovision"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func Prototype() helpers.Transactions {
	return baseHelpers.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		mutate.Transaction,
		nub.Transaction,
		provision.Transaction,
		quash.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	)
}
