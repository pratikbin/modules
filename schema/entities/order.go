package entities

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Order interface {
	GetID() types.ID
	traits.InterChain
	traits.Burnable
	traits.Lockable
	traits.HasImmutables
	GetMakerAddress() sdkTypes.AccAddress
	GetTakerAddress() sdkTypes.AccAddress
	GetMakerAssetAmount() sdkTypes.Dec
	GetMakerAssetData() types.ID
	GetTakerAssetAmount() sdkTypes.Dec
	GetTakerAssetData() types.ID
	GetSalt() types.Height
	SetTakerAddress(sdkTypes.AccAddress) Order
}
