package asset

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, QueryRoute+"/"+"request", nil)
	codec.RegisterConcrete(queryResponse{}, QueryRoute+"/"+"response", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
