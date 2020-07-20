package burn

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper types.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	//message := messageFromInterface(msg)
	//assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(message.OrderID)
	//order := assets.Get(message.OrderID)
	//if order == nil {
	//	return constants.EntityNotFound
	//}
	//if !order.CanBurn(types.NewHeight(context.BlockHeight())) {
	//	return constants.BurnNotAllowed
	//}
	//assets.Remove(order)
	return nil
}

func initializeTransactionKeeper(mapper types.Mapper, externalKeepers []interface{}) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
