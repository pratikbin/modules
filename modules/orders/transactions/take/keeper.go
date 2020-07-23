package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/orders/mapper"
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

type transactionKeeper struct {
	mapper utility.Mapper
}

var _ utility.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	orderID := message.OrderID
	orders := mapper.NewOrders(transactionKeeper.mapper, context).Fetch(orderID)
	order := orders.Get(orderID)
	if order == nil {
		return constants.EntityNotFound
	}

	order = mapper.NewOrder(order.GetID(), order.GetBurn(), order.GetLock(), order.GetImmutables(), order.GetMutables(),
		order.GetMakerAddress(), message.From, order.GetMakerAssetAmount(), order.GetMakerAssetData(), order.GetTakerAssetAmount(),
		order.GetTakerAssetData(), order.GetSalt())
	orders = orders.Mutate(order)
	orders.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper utility.Mapper, externalKeepers []interface{}) utility.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
