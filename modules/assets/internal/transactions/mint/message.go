// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	xprtErrors "github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/assets/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/AssetMantle/modules/utilities/transaction"
)

type message struct {
	From                    sdkTypes.AccAddress  `json:"from" valid:"required~required field from missing"`
	FromID                  types.ID             `json:"fromID" valid:"required~required field fromID missing"`
	ToID                    types.ID             `json:"toID" valid:"required~required field toID missing"`
	ClassificationID        types.ID             `json:"classificationID" valid:"required~required field classificationID missing"`
	ImmutableMetaProperties types.MetaProperties `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing"`
	ImmutableProperties     types.Properties     `json:"immutableProperties" valid:"required~required field immutableProperties missing"`
	MutableMetaProperties   types.MetaProperties `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing"`
	MutableProperties       types.Properties     `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ helpers.Message = message{}

func (message message) Route() string { return module.Name }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(xprtErrors.IncorrectMessage, Error.Error())
	}

	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}
func (message) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, module.Name, message{})
}
func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}
func messagePrototype() helpers.Message {
	return message{}
}
func newMessage(from sdkTypes.AccAddress, fromID types.ID, toID types.ID, classificationID types.ID, immutableMetaProperties types.MetaProperties, immutableProperties types.Properties, mutableMetaProperties types.MetaProperties, mutableProperties types.Properties) sdkTypes.Msg {
	return message{
		From:                    from,
		FromID:                  fromID,
		ToID:                    toID,
		ClassificationID:        classificationID,
		ImmutableMetaProperties: immutableMetaProperties,
		ImmutableProperties:     immutableProperties,
		MutableMetaProperties:   mutableMetaProperties,
		MutableProperties:       mutableProperties,
	}
}
