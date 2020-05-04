package lock

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/asset/mapper"
)

type Keeper interface {
	transact(sdkTypes.Context, Message) error
}

type baseKeeper struct {
	mapper mapper.Mapper
}

func NewKeeper(mapper mapper.Mapper) Keeper {
	return baseKeeper{mapper: mapper}
}

var _ Keeper = (*baseKeeper)(nil)

func (baseKeeper baseKeeper) transact(context sdkTypes.Context, message Message) error {
	asset, err := baseKeeper.mapper.Read(context, mapper.NewAssetAddress(message.Address))
	if err != nil {
		return err
	}
	asset.SetLock(message.Lock)
	return baseKeeper.mapper.Update(context, asset)
}
