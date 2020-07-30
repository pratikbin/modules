package asset

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type queryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context sdkTypes.Context, queryRequest helpers.QueryRequest) helpers.QueryResponse {
	return newQueryResponse(mapper.NewAssets(queryKeeper.mapper, context).Fetch(queryRequestFromInterface(queryRequest).AssetID))
}

func initializeQueryKeeper(mapper helpers.Mapper, _ []interface{}) helpers.QueryKeeper {
	return queryKeeper{mapper: mapper}
}
