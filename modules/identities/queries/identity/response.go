package identity

import (
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type queryResponse struct {
	Identities schema.InterIdentities `json:"identities" valid:"required~Enter the Identities"`
}

var _ utility.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() utility.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(identities schema.InterIdentities) utility.QueryResponse {
	return queryResponse{Identities: identities}
}
