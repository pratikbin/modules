// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package burn

import (
	"context"
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	"github.com/cosmos/cosmos-sdk/types"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	renumerateAuxiliary   helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context types.Context, message helpers.Message) helpers.TransactionResponse {
	_, err := transactionKeeper.Handle(context.Context(), message.(*Message))
	return newTransactionResponse(err)
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*Response, error) {

	fromAddress, err := types.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets := transactionKeeper.mapper.NewCollection(types.UnwrapSDKContext(context)).Fetch(key.NewKey(message.AssetID))

	Mappable := assets.Get(key.NewKey(message.AssetID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound
	}
	asset := Mappable.(documents.Asset)

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), supplement.NewAuxiliaryRequest(asset.GetBurn())))
	if err != nil {
		return nil, err
	}

	if burnHeightMetaProperty := metaProperties.GetProperty(constants.BurnHeightProperty.GetID()); burnHeightMetaProperty != nil {
		burnHeight := burnHeightMetaProperty.(properties.MetaProperty).GetData().(data.HeightData).Get()
		if burnHeight.Compare(baseTypes.NewHeight(types.UnwrapSDKContext(context).BlockHeight())) > 0 {
			return nil, errorConstants.NotAuthorized
		}
	}

	if auxiliaryResponse := transactionKeeper.renumerateAuxiliary.GetKeeper().Help(types.UnwrapSDKContext(context), renumerate.NewAuxiliaryRequest(message.FromID, message.AssetID, types.ZeroDec())); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	assets.Remove(mappable.NewMappable(asset))

	return &Response{}, nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case renumerate.Auxiliary.GetName():
				transactionKeeper.renumerateAuxiliary = value
			case maintain.Auxiliary.GetName():
				transactionKeeper.maintainAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
