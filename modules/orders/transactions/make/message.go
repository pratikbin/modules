package make

import (
	"crypto/sha512"
	"encoding/base64"
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

//TODO make private
type Message struct {
	From             sdkTypes.AccAddress
	Properties       types.Properties
	Lock             types.Height
	Burn             types.Height
	TakerAddress     sdkTypes.AccAddress
	MakerAssetAmount sdkTypes.Dec
	MakerAssetData   types.ID
	MakerAssetType   types.ID
	TakerAssetAmount sdkTypes.Dec
	TakerAssetData   types.ID
	TakerAssetType   types.ID
}

var _ sdkTypes.Msg = Message{}

func (message Message) Route() string { return Transaction.GetModuleName() }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}

func (message Message) GenerateHash(salt types.Height) types.ID {
	hash := sha512.New()
	bz := []byte(message.From.String() + message.TakerAddress.String() +
		message.MakerAssetAmount.String() + message.MakerAssetData.String() +
		message.TakerAssetAmount.String() + message.TakerAssetData.String() + string(salt.Get()))
	hash.Write(bz)

	sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return base.NewID(sha)
}

func newMessage(from sdkTypes.AccAddress, properties types.Properties, lock types.Height, burn types.Height,
	takerAddress sdkTypes.AccAddress, makerAssetAmount sdkTypes.Dec, makerAssetData types.ID, makerAssetType types.ID,
	takerAssetAmount sdkTypes.Dec, takerAssetData types.ID, takerAssetType types.ID) sdkTypes.Msg {
	return Message{
		From:             from,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
		TakerAddress:     takerAddress,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		MakerAssetType:   makerAssetType,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
		TakerAssetType:   takerAssetType,
	}
}
