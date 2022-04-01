// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package split

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
	"splits",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	flags.SplitID,
)
