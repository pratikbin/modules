// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transactions

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/module/transactions/define"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/deputize"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/issue"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/nub"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/provision"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/revoke"
	"github.com/AssetMantle/modules/modules/identities/module/transactions/unprovision"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	require.Equal(t, Prototype().Get("unprovision").GetName(), baseHelpers.NewTransactions(
		define.Transaction,
		deputize.Transaction,
		issue.Transaction,
		nub.Transaction,
		provision.Transaction,
		revoke.Transaction,
		unprovision.Transaction,
	).Get("unprovision").GetName())
}
