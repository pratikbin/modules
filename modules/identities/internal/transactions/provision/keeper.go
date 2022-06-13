// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package provision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))

	identity := identities.Get(key.FromID(identityID)).(mappables.Identity)
	if identity == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	if !utilities.IsProvisioned(identity, message.From) {
		return newTransactionResponse(errors.NotAuthorized)
	}

	if utilities.IsProvisioned(identity, message.To) {
		return newTransactionResponse(errors.EntityAlreadyExists)
	}

	identities.Mutate(utilities.ProvisionAddress(identity, message.To))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
