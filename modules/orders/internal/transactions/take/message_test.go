package take

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/utilities/transaction"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Take_Message(t *testing.T) {

	testFromID := base.NewID("fromID")
	testOrderID := base.NewID("orderID")
	testTakerOwnableSplit := sdkTypes.NewDec(2)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testMessage := newMessage(fromAccAddress, testFromID, testTakerOwnableSplit, testOrderID)
	require.Equal(t, message{From: fromAccAddress, FromID: testFromID, TakerOwnableSplit: testTakerOwnableSplit, OrderID: testOrderID}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())

}
