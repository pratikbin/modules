// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package renumerate

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp"
	"reflect"
	"testing"

	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/renumerate"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	maintainAuxiliary     helpers.Auxiliary
	renumerateAuxiliary   helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
)

type TestKeepers struct {
	RenumerateKeeper helpers.TransactionKeeper
}

func createTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	maintainAuxiliary = maintain.AuxiliaryMock.Initialize(Mapper, Parameters)
	renumerateAuxiliary = renumerate.AuxiliaryMock.Initialize(Mapper, Parameters)
	supplementAuxiliary = supplement.AuxiliaryMock.Initialize(Mapper, Parameters)
	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)

	keepers := TestKeepers{
		RenumerateKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, Mapper, Parameters := createTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		maintainAuxiliary     helpers.Auxiliary
		renumerateAuxiliary   helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		in1         helpers.Parameters
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve", fields{Mapper, maintainAuxiliary, renumerateAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, maintainAuxiliary, renumerateAuxiliary, supplementAuxiliary, authenticateAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				renumerateAuxiliary:   tt.fields.renumerateAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, _ := createTestInput(t)
	immutableProperties := baseLists.NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")), baseProperties.NewMetaProperty(baseIDs.NewStringID("supply"), baseData.NewDecData(sdkTypes.NewDec(1)))}...)
	immutables := baseQualified.NewImmutables(immutableProperties)
	mutableProperties := baseLists.NewPropertyList([]properties.Property{baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList()))}...)
	mutables := baseQualified.NewMutables(mutableProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testAsset := base.NewAsset(classificationID, immutables, mutables)
	testAssetID := baseIDs.NewAssetID(classificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromID := baseIDs.NewIdentityID(classificationID, immutables)
	keepers.RenumerateKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testAsset))
	type fields struct {
		mapper                helpers.Mapper
		maintainAuxiliary     helpers.Auxiliary
		renumerateAuxiliary   helpers.Auxiliary
		supplementAuxiliary   helpers.Auxiliary
		authenticateAuxiliary helpers.Auxiliary
	}
	type args struct {
		context sdkTypes.Context
		msg     sdkTypes.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		{"+ve", fields{Mapper, maintainAuxiliary, renumerateAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, fromID, testAssetID)}, newTransactionResponse(nil)},
		{"+ve", fields{Mapper, maintainAuxiliary, renumerateAuxiliary, supplementAuxiliary, authenticateAuxiliary}, args{context, newMessage(fromAccAddress, fromID, baseIDs.PrototypeAssetID())}, newTransactionResponse(errorConstants.EntityNotFound)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				renumerateAuxiliary:   tt.fields.renumerateAuxiliary,
				supplementAuxiliary:   tt.fields.supplementAuxiliary,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
