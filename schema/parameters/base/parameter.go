// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/parameters"
)

type parameter struct {
	ID        ids.ID    `json:"id"`
	Data      data.Data `json:"data"`
	validator func(interface{}) error
}

var _ parameters.Parameter = (*parameter)(nil)

func (parameter parameter) String() string {
	return parameter.ID.String()
}
func (parameter parameter) Equal(compareParameter parameters.Parameter) bool {
	if compareParameter != nil && parameter.ID.Compare(compareParameter.GetID()) == 0 && parameter.Data.GetType().Compare(compareParameter.GetData().GetType()) == 0 && parameter.Data.Compare(compareParameter.GetData()) == 0 {
		return true
	}

	return false
}
func (parameter parameter) Validate() error {
	return parameter.validator(parameter)
}
func (parameter parameter) GetID() ids.ID {
	return parameter.ID
}
func (parameter parameter) GetData() data.Data {
	return parameter.Data
}
func (parameter parameter) GetValidator() func(interface{}) error {
	return parameter.validator
}
func (parameter parameter) Mutate(data data.Data) parameters.Parameter {
	parameter.Data = data
	return parameter
}

func NewParameter(id ids.ID, data data.Data, validator func(interface{}) error) parameters.Parameter {
	return parameter{
		ID:        id,
		Data:      data,
		validator: validator,
	}
}
