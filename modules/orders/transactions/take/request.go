/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package take

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq           rest.BaseReq `json:"baseReq"`
	FromID            string       `json:"fromID" valid:"required~required field fromID missing"`
	TakerOwnableSplit string       `json:"takerOwnableSplit" valid:"required~required field takerOwnableSplit missing"`
	OrderID           string       `json:"orderID" valid:"required~required field orderID missing"`
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
		cliCommand.ReadString(flags.TakerOwnableSplit),
		cliCommand.ReadString(flags.OrderID),
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

	takerOwnableSplit, Error := sdkTypes.NewDecFromStr(transactionRequest.TakerOwnableSplit)
	if Error != nil {
		return nil, Error
	}
	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		takerOwnableSplit,
		base.NewID(transactionRequest.OrderID),
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, takerOwnableSplit string, orderID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:           baseReq,
		FromID:            fromID,
		TakerOwnableSplit: takerOwnableSplit,
		OrderID:           orderID,
	}
}
