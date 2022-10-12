// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"fmt"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/scrub"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	base2 "github.com/AssetMantle/modules/schema/helpers/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	base3 "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/test/schema/helpers/base"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/stretchr/testify/require"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

func createTestInput(t *testing.T) (helpers.Mapper, helpers.Auxiliary, helpers.Auxiliary, helpers.Auxiliary, helpers.Auxiliary, helpers.Keeper, helpers.Parameters, types.Context) {
	testContext, storeKey, paramsTransientStoreKeys := base.SetupTest(t)
	testMapper := base2.NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	paramsStoreKey := types.NewKVStoreKey("testParams")
	paramsKeeper := params.NewKeeper(
		codec.Cdc,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	testParameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, types.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, types.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	testDefineAuxiliary := define.AuxiliaryMock.Initialize(testMapper, testParameters)
	testScrubAuxiliary := scrub.AuxiliaryMock.Initialize(testMapper, testParameters)
	testSuperAuxiliary := super.AuxiliaryMock.Initialize(testMapper, testParameters)
	testVerifyAuxiliary := verify.AuxiliaryMock.Initialize(testMapper, testParameters)
	testKeepers := keeperPrototype().Initialize(testMapper, testParameters, []interface{}{testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary}).(helpers.TransactionKeeper)
	return testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary, testKeepers, testParameters, testContext
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
	testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary, _, testParameters, _ := createTestInput(t)
	type fields struct {
		mapper          helpers.Mapper
		defineAuxiliary helpers.Auxiliary
		scrubAuxiliary  helpers.Auxiliary
		superAuxiliary  helpers.Auxiliary
		verifyAuxiliary helpers.Auxiliary
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
		{"+ve", fields{testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary}, args{testMapper, testParameters, []interface{}{testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary}}, transactionKeeper{testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:          tt.fields.mapper,
				defineAuxiliary: tt.fields.defineAuxiliary,
				scrubAuxiliary:  tt.fields.scrubAuxiliary,
				superAuxiliary:  tt.fields.superAuxiliary,
				verifyAuxiliary: tt.fields.verifyAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !compare(got, tt.want) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	testImmutables := baseQualified.NewImmutables(base3.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	testMutables := baseQualified.NewMutables(base3.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	testClassificationID := baseIDs.NewClassificationID(testImmutables, testMutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, testImmutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	immutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultImmutableMeta1:S|defaultImmutableMeta1")
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaPropertyList("defaultMutableMeta1:S|defaultMutableMeta1")
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties("defaultMutable1:S|defaultMutable1")
	require.Equal(t, nil, err)
	testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary, _, _, testContext := createTestInput(t)
	type fields struct {
		mapper          helpers.Mapper
		defineAuxiliary helpers.Auxiliary
		scrubAuxiliary  helpers.Auxiliary
		superAuxiliary  helpers.Auxiliary
		verifyAuxiliary helpers.Auxiliary
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
		{"+ve", fields{testMapper, testDefineAuxiliary, testScrubAuxiliary, testSuperAuxiliary, testVerifyAuxiliary}, args{testContext, newMessage(fromAccAddress, testIdentityID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)}, newTransactionResponse(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:          tt.fields.mapper,
				defineAuxiliary: tt.fields.defineAuxiliary,
				scrubAuxiliary:  tt.fields.scrubAuxiliary,
				superAuxiliary:  tt.fields.superAuxiliary,
				verifyAuxiliary: tt.fields.verifyAuxiliary,
			}
			if got := transactionKeeper.Transact(tt.args.context, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compare(x, y interface{}) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	//fmt.Println(v1.Field(4))
	for i := 0; i < v1.NumField(); i++ {
		if !reflect.DeepEqual(fmt.Sprint(v1.Field(i)), fmt.Sprint(v2.Field(i))) {
			fmt.Println(v1.Field(i), v2.Field(i))
			return false
		}
	}
	return true
}
