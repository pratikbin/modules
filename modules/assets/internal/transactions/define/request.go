/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package define

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type transactionRequest struct {
	BaseReq             rest.BaseReq `json:"baseReq"`
	FromID              string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ImmutableMetaTraits string       `json:"immutableMetaTraits" valid:"required~required field immutableMetaTraits missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableTraits     string       `json:"immutableTraits" valid:"required~required field immutableTraits missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaTraits   string       `json:"mutableMetaTraits" valid:"required~required field mutableMetaTraits missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableTraits       string       `json:"mutableTraits" valid:"required~required field mutableTraits missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ImmutableMetaTraits),
		cliCommand.ReadString(flags.ImmutableTraits),
		cliCommand.ReadString(flags.MutableMetaTraits),
		cliCommand.ReadString(flags.MutableTraits),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	immutableMetaTraits, Error := base.ReadMetaProperties(transactionRequest.ImmutableMetaTraits)
	if Error != nil {
		return nil, Error
	}

	immutableTraits, Error := base.ReadProperties(transactionRequest.ImmutableTraits)
	if Error != nil {
		return nil, Error
	}

	mutableMetaTraits, Error := base.ReadMetaProperties(transactionRequest.MutableMetaTraits)
	if Error != nil {
		return nil, Error
	}

	mutableTraits, Error := base.ReadProperties(transactionRequest.MutableTraits)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		immutableMetaTraits,
		immutableTraits,
		mutableMetaTraits,
		mutableTraits,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, immutableMetaTraits string, immutableTraits string, mutableMetaTraits string, mutableTraits string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:             baseReq,
		FromID:              fromID,
		ImmutableMetaTraits: immutableMetaTraits,
		ImmutableTraits:     immutableTraits,
		MutableMetaTraits:   mutableMetaTraits,
		MutableTraits:       mutableTraits,
	}
}
