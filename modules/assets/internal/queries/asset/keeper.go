/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(queryKeeper.mapper.NewCollection(context).Fetch(key.New(queryRequestFromInterface(queryRequest).AssetID)), nil)
}

func initializeQueryKeeper(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
