/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package block

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/scrub"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/transfer"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type block struct {
	mapper              helpers.Mapper
	parameters          helpers.Parameters
	supplementAuxiliary helpers.Auxiliary
	transferAuxiliary   helpers.Auxiliary
	scrubAuxiliary      helpers.Auxiliary
}

var _ helpers.Block = (*block)(nil)

func (block block) Begin(_ sdkTypes.Context, _ abciTypes.RequestBeginBlock) {

}

func (block block) End(context sdkTypes.Context, _ abciTypes.RequestEndBlock) {
	executeOrders := make(map[types.ID]bool)
	orders := block.mapper.NewCollection(context)

	orders.Iterate(key.New(base.NewID("")), func(order helpers.Mappable) bool {
		id1 := key.NewOrderID(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetMakerOwnableID(), order.(mappables.Order).GetTakerOwnableID(), base.NewID(""), base.NewID(""), base.NewID(""), base.NewImmutables(base.NewProperties()))
		id2 := key.NewOrderID(order.(mappables.Order).GetClassificationID(), order.(mappables.Order).GetTakerOwnableID(), order.(mappables.Order).GetMakerOwnableID(), base.NewID(""), base.NewID(""), base.NewID(""), base.NewImmutables(base.NewProperties()))
		if !executeOrders[id1] && !executeOrders[id2] {
			executeOrders[id1] = true
		}
		return false
	})

	for partialOrderID := range executeOrders {
		nextPartialOrderID := false

		orders.Iterate(key.New(partialOrderID), func(orderMappable helpers.Mappable) bool {
			orders.Iterate(
				key.New(key.NewOrderID(orderMappable.(mappables.Order).GetClassificationID(), orderMappable.(mappables.Order).GetTakerOwnableID(), orderMappable.(mappables.Order).GetMakerOwnableID(), base.NewID(""), base.NewID(""), base.NewID(""), base.NewImmutables(base.NewProperties()))),
				func(executableMappableOrder helpers.Mappable) bool {

					var leftOrder mappables.Order
					var rightOrder mappables.Order

					orderHeight, Error := orderMappable.(mappables.Order).GetCreation().(types.MetaProperty).GetMetaFact().GetData().AsHeight()
					if Error != nil {
						panic(Error)
					}

					executableOrderHeight, Error := executableMappableOrder.(mappables.Order).GetCreation().(types.MetaProperty).GetMetaFact().GetData().AsHeight()
					if Error != nil {
						panic(Error)
					}

					switch {
					case orderHeight.IsGreaterThan(executableOrderHeight):
						leftOrder = orderMappable.(mappables.Order)
						rightOrder = executableMappableOrder.(mappables.Order)
					case executableOrderHeight.IsGreaterThan(orderHeight):
						leftOrder = executableMappableOrder.(mappables.Order)
						rightOrder = orderMappable.(mappables.Order)
					default:
						// TODO
						leftOrder = orderMappable.(mappables.Order)
						rightOrder = executableMappableOrder.(mappables.Order)
					}

					leftOrderExchangeRate, Error := leftOrder.GetExchangeRate().(types.MetaProperty).GetMetaFact().GetData().AsDec()
					if Error != nil {
						panic(Error)
					}

					leftOrderMakerOwnableSplit, Error := getMakerOwnableSplit(context, leftOrder, block.supplementAuxiliary)
					if Error != nil {
						panic(Error)
					}

					rightOrderExchangeRate, Error := rightOrder.GetExchangeRate().(types.MetaProperty).GetMetaFact().GetData().AsDec()
					if Error != nil {
						panic(Error)
					}

					rightOrderMakerOwnableSplit, Error := getMakerOwnableSplit(context, rightOrder, block.supplementAuxiliary)
					if Error != nil {
						panic(Error)
					}

					rightOrderTakerOwnableSplitDemanded := rightOrderExchangeRate.Mul(rightOrderMakerOwnableSplit).MulTruncate(sdkTypes.SmallestDec())

					if leftOrderExchangeRate.MulTruncate(rightOrderExchangeRate).LTE(sdkTypes.OneDec()) {
						switch {
						case leftOrderMakerOwnableSplit.GT(rightOrderTakerOwnableSplitDemanded):
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), rightOrderTakerOwnableSplitDemanded)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							mutableProperties, Error := scrub.GetPropertiesFromResponse(block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(leftOrderMakerOwnableSplit.Sub(rightOrderTakerOwnableSplitDemanded)))))))
							if Error != nil {
								panic(Error)
							}

							orders.Mutate(mappable.NewOrder(leftOrder.GetID(), leftOrder.GetImmutables(), leftOrder.GetMutables().Mutate(mutableProperties.GetList()...)))
							orders.Remove(rightOrder)

							if executableOrderHeight.IsGreaterThan(orderHeight) {
								return true
							}
						case leftOrderMakerOwnableSplit.LT(rightOrderTakerOwnableSplitDemanded):
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), leftOrderMakerOwnableSplit.QuoTruncate(rightOrderExchangeRate))); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							mutableProperties, Error := scrub.GetPropertiesFromResponse(block.scrubAuxiliary.GetKeeper().Help(context, scrub.NewAuxiliaryRequest(base.NewMetaProperty(base.NewID(properties.MakerOwnableSplit), base.NewMetaFact(base.NewDecData(rightOrderMakerOwnableSplit.Sub(leftOrderMakerOwnableSplit.QuoTruncate(rightOrderExchangeRate))))))))
							if Error != nil {
								panic(Error)
							}

							orders.Mutate(mappable.NewOrder(rightOrder.GetID(), rightOrder.GetImmutables(), rightOrder.GetMutables().Mutate(mutableProperties.GetList()...)))
							orders.Remove(leftOrder)

							if orderHeight.IsGreaterThan(executableOrderHeight) || orderHeight.Equals(executableOrderHeight) {
								return true
							}
						default:
							// case leftOrderMakerOwnableSplit.Equal(rightOrderTakerOwnableSplitDemanded):
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), leftOrder.GetMakerID(), leftOrder.GetTakerOwnableID(), rightOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}
							if auxiliaryResponse := block.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(base.NewID(module.Name), rightOrder.GetMakerID(), leftOrder.GetMakerOwnableID(), leftOrderMakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
								panic(auxiliaryResponse.GetError())
							}

							orders.Remove(rightOrder)
							orders.Remove(leftOrder)
							return true
						}
					} else {
						nextPartialOrderID = true
						return true
					}

					return false
				},
			)

			return nextPartialOrderID
		})

		if nextPartialOrderID {
			continue
		}
	}
}

func (block block) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaryKeepers ...interface{}) helpers.Block {
	block.mapper, block.parameters = mapper, parameters

	for _, auxiliaryKeeper := range auxiliaryKeepers {
		switch value := auxiliaryKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				block.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				block.transferAuxiliary = value
			case scrub.Auxiliary.GetName():
				block.scrubAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return block
}

func getMakerOwnableSplit(context sdkTypes.Context, order mappables.Order, supplementAuxiliary helpers.Auxiliary) (sdkTypes.Dec, error) {
	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(order.(mappables.Order).GetMakerOwnableSplit())))
	if Error != nil {
		panic(Error)
	}

	makerOwnableSplit, Error := metaProperties.GetMetaProperty(base.NewID(properties.MakerOwnableSplit)).GetMetaFact().GetData().AsDec()

	return makerOwnableSplit, Error
}
