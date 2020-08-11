/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package genesis

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"gopkg.in/validator.v2"
)

type genesisState struct {
	ClassificationList []mappables.Classification
}

var _ helpers.GenesisState = (*genesisState)(nil)

func (genesisState genesisState) Default() helpers.GenesisState {
	return genesisState
}


func (genesisState genesisState) Validate(sdkTypes.Context) error {

	for _, classification := range genesisState.ClassificationList {
		if errs := validator.Validate(classification); errs != nil {
			return errs
		}

		if classification.GetID() == nil { return constants.EntityNotFound}
		traits := classification.GetTraits().GetList()
		for _, trait := range traits{
			if trait.GetID() != trait.GetProperty().GetID() {
				return constants.IncorrectMessage
			}
		}
	}
	return nil
}


func (genesisState genesisState) Initialize(ctx sdkTypes.Context, mapper helpers.Mapper) {

	for _, cls := range genesisState.ClassificationList {
		mapper.Create(ctx, cls)
	}
}

func (genesisState genesisState) Export(context sdkTypes.Context, mapper helpers.Mapper) helpers.GenesisState {
	assetsID := base.NewID("")

	appendableAssetList := func(mappable traits.Mappable) bool {
		genesisState.ClassificationList = append(genesisState.ClassificationList, mappable.(mappables.Classification))
		return false
	}
	mapper.Iterate(context, assetsID, appendableAssetList)
	return genesisState
}

func (genesisState genesisState) Marshall() []byte {
	return packageCodec.MustMarshalJSON(genesisState)
}
func (genesisState genesisState) Unmarshall(byte []byte) helpers.GenesisState {
	if Error := packageCodec.UnmarshalJSON(byte, &genesisState); Error != nil {
		return nil
	}
	return genesisState
}

func newGenesisState(classificationList []mappables.Classification) helpers.GenesisState {
	return genesisState{
		ClassificationList: classificationList,
	}
}

var State = newGenesisState([]mappables.Classification{})
