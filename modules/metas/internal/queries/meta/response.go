/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package meta

import (
	"fmt"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse QueryResponse) GetError() error {
	return fmt.Errorf(queryResponse.Error)
}
func (queryResponse QueryResponse) LegacyAminoEncode() ([]byte, error) {
	return common.LegacyAminoCodec.MarshalJSON(queryResponse)
}
func (queryResponse QueryResponse) LegacyAminoDecode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.LegacyAminoCodec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return &queryResponse, nil
}
func (queryResponse QueryResponse) Encode() ([]byte, error) {
	return common.JSONCodec.MarshalJSON(&queryResponse)
}

func (queryResponse QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if Error := common.JSONCodec.UnmarshalJSON(bytes, &queryResponse); Error != nil {
		return nil, Error
	}

	return &queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection, error error) QueryResponse {
	success := true
	if error != nil {
		success = false
	}

	return QueryResponse{
		Success: success,
		Error:   error.Error(),
		List:    collection.GetList(),
	}
}
