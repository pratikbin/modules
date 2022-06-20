// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	scrubAuxiliary      helpers.Auxiliary
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(constants2.EntityNotFound)
	}

	orderID := message.OrderID
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(orderID))
	order := orders.Get(key.FromID(orderID))

	if order == nil {
		return newTransactionResponse(constants2.EntityNotFound)
	}

	metaProperties, err := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetTakerID(), order.(mappables.Order).GetMakerOwnableSplit())))
	if err != nil {
		newTransactionResponse(err)
	}

	if takerIDProperty := metaProperties.GetMetaProperty(constants.TakerIDProperty); takerIDProperty != nil {
		takerID := takerIDProperty.GetData().(data.IDData).Get()
		if takerID.Compare(baseIDs.NewID("")) != 0 && takerID.Compare(message.FromID) != 0 {
			return newTransactionResponse(constants2.NotAuthorized)
		}
	}

	exchangeRate := order.(mappables.Order).GetExchangeRate().GetData().(data.DecData).Get()

	makerOwnableSplitProperty := metaProperties.GetMetaProperty(constants.MakerOwnableSplitProperty)
	if makerOwnableSplitProperty == nil {
		return newTransactionResponse(constants2.MetaDataError)
	}

	makerOwnableSplit := makerOwnableSplitProperty.GetData().(data.DecData).Get()

	makerReceiveTakerOwnableSplit := makerOwnableSplit.MulTruncate(exchangeRate).MulTruncate(sdkTypes.SmallestDec())
	takerReceiveMakerOwnableSplit := message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(exchangeRate)

	switch updatedMakerOwnableSplit := makerOwnableSplit.Sub(takerReceiveMakerOwnableSplit); {
	case updatedMakerOwnableSplit.Equal(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return newTransactionResponse(constants2.InsufficientBalance)
		}

		orders.Remove(order)
	case updatedMakerOwnableSplit.LT(sdkTypes.ZeroDec()):
		if message.TakerOwnableSplit.LT(makerReceiveTakerOwnableSplit) {
			return newTransactionResponse(constants2.InsufficientBalance)
		}

		takerReceiveMakerOwnableSplit = makerOwnableSplit

		orders.Remove(order)
	default:
		makerReceiveTakerOwnableSplit = message.TakerOwnableSplit
		mutableProperties, err := scrub.GetPropertiesFromResponse(transactionKeeper.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(constants.MakerOwnableSplitProperty, baseData.NewDecData(updatedMakerOwnableSplit)))))

		if err != nil {
			return newTransactionResponse(err)
		}

		order = mappable.NewOrder(orderID, order.(mappables.Order).GetImmutablePropertyList(), order.(mappables.Order).GetImmutablePropertyList().Mutate(mutableProperties.GetList()...))
		orders.Mutate(order)
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, order.(mappables.Order).GetMakerID(), order.(mappables.Order).GetTakerOwnableID(), makerReceiveTakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(baseIDs.NewID(module.Name), message.FromID, order.(mappables.Order).GetMakerOwnableID(), takerReceiveMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case scrub.Auxiliary.GetName():
				transactionKeeper.scrubAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		default:
			panic(constants2.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
