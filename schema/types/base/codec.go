// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(codec, height{})
	codecUtilities.RegisterModuleConcrete(codec, signature{})
	codecUtilities.RegisterModuleConcrete(codec, split{})
}
