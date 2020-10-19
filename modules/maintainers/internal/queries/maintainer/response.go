/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainer

import (
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryResponse struct {
	Success     bool               `json:"success"`
	Error       error              `json:"error"`
	Maintainers helpers.Collection `json:"maintainers" valid:"required~required field maintainers missing"`
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

func newQueryResponse(maintainers helpers.Collection, error error) helpers.QueryResponse {
	success := true
	if error != nil {
		success = false
	}
	return queryResponse{
		Success:     success,
		Error:       error,
		Maintainers: maintainers,
	}
}
