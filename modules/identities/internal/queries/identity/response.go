/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryResponse struct {
	Success    bool               `json:"success"`
	Error      error              `json:"error"`
	Identities helpers.Collection `json:"identities" valid:"required~required field identities missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func (queryResponse queryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse queryResponse) GetError() error {
	return queryResponse.Error
}
func (queryResponse queryResponse) Encode() ([]byte, error) {
	return module.Codec.MarshalJSON(queryResponse)
}

func (queryResponse queryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := module.Codec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}
	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(identities helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}
	return queryResponse{
		Success:    success,
		Error:      error,
		Identities: identities,
	}
}
