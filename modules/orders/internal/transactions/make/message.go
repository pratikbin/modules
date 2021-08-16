/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
)

var _ helpers.Message = &Message{}

func (message Message) Route() string { return module.Name }
func (message Message) Type() string  { return Transaction.GetName() }
func (message Message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	if message.MakerOwnableID.Compare(message.TakerOwnableID) == 0 {
		return errors.Wrap(xprtErrors.IncorrectMessage, "")
	}

	if message.TakerOwnableSplit.LTE(sdkTypes.ZeroDec()) || message.MakerOwnableSplit.LTE(sdkTypes.ZeroDec()) {
		return errors.Wrap(xprtErrors.IncorrectMessage, "")
	}

	return nil
}
func (message Message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(message))
}
func (message Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From.AsSDKTypesAccAddress()}
}
func (Message) RegisterLegacyAminoCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Message{})
}
func messageFromInterface(msg sdkTypes.Msg) Message {
	switch value := msg.(type) {
	case Message:
		return value
	default:
		return Message{}
	}
}
func messagePrototype() helpers.Message {
	return Message{}
}

func newMessage(from sdkTypes.AccAddress, fromID types.ID, classificationID types.ID, makerOwnableID types.ID, takerOwnableID types.ID, expiresIn types.Height, makerOwnableSplit sdkTypes.Dec, takerOwnableSplit sdkTypes.Dec, immutableMetaProperties types.MetaProperties, immutableProperties types.Properties, mutableMetaProperties types.MetaProperties, mutableProperties types.Properties) sdkTypes.Msg {
	return Message{
		From:                    base.NewAccAddressFromSDKTypesAccAddress(from),
		FromID:                  fromID,
		ClassificationID:        classificationID,
		MakerOwnableID:          makerOwnableID,
		TakerOwnableID:          takerOwnableID,
		ExpiresIn:               expiresIn,
		MakerOwnableSplit:       makerOwnableSplit,
		TakerOwnableSplit:       takerOwnableSplit,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
