package mint

import (
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Request request
var Transaction = types.NewTransaction(
	constants.ModuleName,
	constants.MintTransaction,
	constants.MintTransactionShort,
	constants.MintTransactionLong,
	requestPrototype,
	[]types.CLIFlag{
		constants.Properties,
		constants.ChainID,
		constants.MaintainersID,
		constants.ClassificationID,
		constants.Lock,
		constants.Burn},
)
