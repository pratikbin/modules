package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type Keeper interface {
	transact(sdkTypes.Context, message) error
}

type baseKeeper struct {
	mapper mapper.Mapper
}

func NewKeeper(mapper mapper.Mapper) Keeper {
	return baseKeeper{mapper: mapper}
}

var _ Keeper = (*baseKeeper)(nil)

func (baseKeeper baseKeeper) transact(context sdkTypes.Context, message message) error {
	immutablePropertyList := message.propertyList
	hashID := baseKeeper.mapper.GenerateHashID(immutablePropertyList)
	assetID := baseKeeper.mapper.GenerateAssetID(message.chainID, message.maintainersID, message.classificationID, hashID)
	asset := baseKeeper.mapper.MakeAsset(assetID, &types.BaseProperties{PropertyList: message.propertyList}, message.lock, message.burn)
	return baseKeeper.mapper.New(context).Add(asset)
}
