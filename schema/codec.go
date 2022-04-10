// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	baseTraits "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterInterface((*error)(nil), nil)

	data.RegisterCodec(codec)
	baseData.RegisterCodec(codec)

	helpers.RegisterCodec(codec)

	ids.RegisterCodec(codec)
	baseIDs.RegisterCodec(codec)

	mappables.RegisterCodec(codec)

	traits.RegisterCodec(codec)
	baseTraits.RegisterCodec(codec)

	types.RegisterCodec(codec)
	baseTypes.RegisterCodec(codec)

}
