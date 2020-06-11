package mint

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type request struct {
	baseReq          rest.BaseReq
	chainID          types.BaseID
	classificationID types.BaseID
	maintainersID    types.BaseID
	properties       []types.BaseProperty
	lock             types.BaseHeight
	burn             types.BaseHeight
}

func RestRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		var request request
		if !rest.ReadRESTReq(responseWriter, httpRequest, cliContext.Codec, &request) {
			return
		}

		request.baseReq = request.baseReq.Sanitize()
		if !request.baseReq.ValidateBasic(responseWriter) {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, "")
			return
		}

		_, Error := govalidator.ValidateStruct(request)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		from, Error := sdkTypes.AccAddressFromBech32(request.baseReq.From)
		if Error != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, Error.Error())
			return
		}

		noOfPropertiesSent := len(request.properties)

		if noOfPropertiesSent > constants.MaxTraitCount {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, constants.IncorrectMessageCode.Error())
			return
		}

		var basePropertyList []types.BaseProperty
		for _, baseProperty := range request.properties {
			basePropertyList = append(basePropertyList, baseProperty)
		}
		baseProperties := types.BaseProperties{BasePropertyList: basePropertyList}
		message := Message{
			From:             from,
			ChainID:          request.chainID,
			MaintainersID:    request.maintainersID,
			ClassificationID: request.classificationID,
			Properties:       &baseProperties,
			Lock:             request.lock,
			Burn:             request.burn,
		}
		client.WriteGenerateStdTxResponse(responseWriter, cliContext, request.baseReq, []sdkTypes.Msg{message})
	}
}
