// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package immediate

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/constants/flags"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type transactionRequest struct {
	BaseReq                 rest.BaseReq `json:"baseReq"`
	FromID                  string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ClassificationID        string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
	MakerOwnableID          string       `json:"makerOwnableID" valid:"required~required field makerOwnableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field makerOwnableID"`
	TakerOwnableID          string       `json:"takerOwnableID" valid:"required~required field takerOwnableID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field takerOwnableID"`
	ExpiresIn               int64        `json:"expiresIn" valid:"required~required field expiresIn missing, matches(^[0-9]+$)~invalid field expiresIn"`
	MakerOwnableSplit       string       `json:"makerOwnableSplit" valid:"required~required field makerOwnableSplit missing, matches(^[0-9.]+$)~invalid field makerOwnableSplit"`
	TakerOwnableSplit       string       `json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing, matches(^[0-9.]+$)~invalid field takerOwnableSplit"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Immediate order transaction
// @Description Immediate order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body  transactionRequest true "Request body for immediate order"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/immediate [post]
func (transactionRequest transactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.ClassificationID),
		cliCommand.ReadString(flags.MakerOwnableID),
		cliCommand.ReadString(flags.TakerOwnableID),
		cliCommand.ReadInt64(flags.ExpiresIn),
		cliCommand.ReadString(flags.MakerOwnableSplit),
		cliCommand.ReadString(flags.TakerOwnableSplit),
		cliCommand.ReadString(flags.ImmutableMetaProperties),
		cliCommand.ReadString(flags.ImmutableProperties),
		cliCommand.ReadString(flags.MutableMetaProperties),
		cliCommand.ReadString(flags.MutableProperties),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}

func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if err != nil {
		return nil, err
	}

	makerOwnableSplit, err := sdkTypes.NewDecFromStr(transactionRequest.MakerOwnableSplit)
	if err != nil {
		return nil, err
	}

	takerOwnableSplit, err := sdkTypes.NewDecFromStr(transactionRequest.TakerOwnableSplit)
	if err != nil {
		return nil, err
	}

	immutableMetaProperties, err := base.ReadMetaProperties(transactionRequest.ImmutableMetaProperties)
	if err != nil {
		return nil, err
	}

	immutableProperties, err := base.ReadProperties(transactionRequest.ImmutableProperties)
	if err != nil {
		return nil, err
	}

	mutableMetaProperties, err := base.ReadMetaProperties(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := base.ReadProperties(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ClassificationID),
		base.NewID(transactionRequest.MakerOwnableID),
		base.NewID(transactionRequest.TakerOwnableID),
		base.NewHeight(transactionRequest.ExpiresIn),
		makerOwnableSplit,
		takerOwnableSplit,
		immutableMetaProperties,
		immutableProperties,
		mutableMetaProperties,
		mutableProperties,
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, fromID string, classificationID string, makerOwnableID string, takerOwnableID string, expiresIn int64, makerOwnableSplit, takerOwnableSplit string, immutableMetaProperties string, immutableProperties string, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:                 baseReq,
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
