package base

import (
	"github.com/cosmos/cosmos-sdk/codec"

	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

func RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(codec, metaProperty{})
	codecUtilities.RegisterModuleConcrete(codec, mesaProperty{})
}
