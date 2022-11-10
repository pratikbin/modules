// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"fmt"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/member"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/identities/internal/mappable"
	"github.com/AssetMantle/modules/modules/identities/internal/parameters"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	base2 "github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base3 "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

var (
	authenticateAuxiliary helpers.Auxiliary
	maintainAuxiliary     helpers.Auxiliary
	memberAuxiliary       helpers.Auxiliary
)

type TestKeepers struct {
	MutateKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (types.Context, TestKeepers, helpers.Mapper, helpers.Parameters) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	types.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := types.NewKVStoreKey("test")
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := types.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	authenticateAuxiliary = authenticate.AuxiliaryMock.Initialize(Mapper, Parameters)
	maintainAuxiliary = maintain.AuxiliaryMock.Initialize(Mapper, Parameters)
	memberAuxiliary = member.AuxiliaryMock.Initialize(Mapper, Parameters)

	context := types.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MutateKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.TransactionKeeper),
	}

	return context, keepers, Mapper, Parameters
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		// TODO: Add test cases.
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
	_, _, Mapper, Parameters := CreateTestInput(t)
	type fields struct {
		mapper                helpers.Mapper
		authenticateAuxiliary helpers.Auxiliary
		maintainAuxiliary     helpers.Auxiliary
		memberAuxiliary       helpers.Auxiliary
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
		// TODO: Add test cases.
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, authenticateAuxiliary, maintainAuxiliary, memberAuxiliary}, args{Mapper, Parameters, []interface{}{}}, transactionKeeper{Mapper, authenticateAuxiliary, maintainAuxiliary, memberAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				memberAuxiliary:       tt.fields.memberAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	context, keepers, Mapper, _ := CreateTestInput(t)
	mutableProperties := base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(base.NewDataList())))
	immutableProperties := base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(base.NewDataList())))
	toMutateMetaProperties := base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(base.NewDataList(base.NewStringData("Test")))))
	immutables := baseQualified.NewImmutables(immutableProperties)
	mutables := baseQualified.NewMutables(mutableProperties)
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity := base2.NewIdentity(testClassificationID, immutables, mutables)
	testIdentity.ProvisionAddress([]types.AccAddress{fromAccAddress}...)
	keepers.MutateKeeper.(transactionKeeper).mapper.NewCollection(context).Add(mappable.NewMappable(testIdentity))
	type fields struct {
		mapper                helpers.Mapper
		authenticateAuxiliary helpers.Auxiliary
		maintainAuxiliary     helpers.Auxiliary
		memberAuxiliary       helpers.Auxiliary
	}
	type args struct {
		context types.Context
		msg     types.Msg
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.TransactionResponse
	}{
		// TODO: Add test cases.
		{"+ve", fields{Mapper, authenticateAuxiliary, maintainAuxiliary, memberAuxiliary}, args{context, newMessage(fromAccAddress, testFromID, testFromID, mutableProperties, toMutateMetaProperties)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:                tt.fields.mapper,
				authenticateAuxiliary: tt.fields.authenticateAuxiliary,
				maintainAuxiliary:     tt.fields.maintainAuxiliary,
				memberAuxiliary:       tt.fields.memberAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}
