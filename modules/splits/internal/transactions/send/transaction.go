// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	"send",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,
	flags.FromID,
	flags.ToID,
	flags.OwnableID,
	flags.Value,
)
