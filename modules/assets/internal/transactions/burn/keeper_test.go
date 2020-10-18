/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/metas/auxiliaries/supplement"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	AssetsKeeper helpers.TransactionKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.New()
	assets.Prototype.RegisterCodec(Codec)
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeTransient, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	burn.AuxiliaryMock.Initialize(mapper, Parameters)
	supplement.AuxiliaryMock.Initialize(mapper, Parameters)
	verify.AuxiliaryMock.Initialize(mapper, Parameters)
	keepers := TestKeepers{
		AssetsKeeper: keeperPrototype().Initialize(mapper, Parameters,
			[]interface{}{burn.AuxiliaryMock, supplement.AuxiliaryMock,
				verify.AuxiliaryMock}).(helpers.TransactionKeeper),
	}

	return context, keepers
}

func Test_transactionKeeper_Transact(t *testing.T) {

	ctx, keepers := CreateTestInput(t)
	ctx = ctx.WithBlockHeight(2)
	immutableTraits, Error := base.ReadProperties("defaultImmutable1:S|defaultImmutable1")
	require.Equal(t, nil, Error)
	mutableTraits, Error := base.ReadProperties("burn:S|100")
	require.Equal(t, nil, Error)
	supplementError, Error := base.ReadMetaProperties("supplementError:S|mockError")
	require.Equal(t, nil, Error)
	defaultAddr := sdkTypes.AccAddress("addr")
	verifyMockErrorAddress := sdkTypes.AccAddress("verifyError")
	defaultIdentityID := base.NewID("fromIdentityID")
	burnMockErrorIdentity := base.NewID("burnError")
	classificationID := base.NewID("ClassificationID")
	assetID := key.NewAssetID(classificationID, base.NewImmutables(immutableTraits))
	assetID2 := key.NewAssetID(base.NewID("ClassificationID2"), base.NewImmutables(immutableTraits))
	assetID3 := key.NewAssetID(base.NewID("ClassificationID3"), base.NewImmutables(immutableTraits))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID,
		base.NewImmutables(immutableTraits), base.NewMutables(mutableTraits)))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID2,
		base.NewImmutables(immutableTraits), base.NewMutables(supplementError)))
	keepers.AssetsKeeper.(transactionKeeper).mapper.NewCollection(ctx).Add(mappable.NewAsset(assetID3,
		base.NewImmutables(immutableTraits), base.NewMutables(mutableTraits)))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newTransactionResponse(nil)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - verify identity mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(verifyMockErrorAddress, defaultIdentityID, assetID)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - unMinted asset", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.EntityNotFound)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, base.NewID(""))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - supplement mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, defaultIdentityID, assetID2)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase - burn mock error", func(t *testing.T) {
		t.Parallel()
		want := newTransactionResponse(errors.MockError)
		if got := keepers.AssetsKeeper.Transact(ctx, newMessage(defaultAddr, burnMockErrorIdentity, assetID3)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})
	t.Run("NegativeCase - burn height error", func(t *testing.T) {
		ctx2 := ctx.WithBlockHeight(-20)
		want := newTransactionResponse(errors.NotAuthorized)
		if got := keepers.AssetsKeeper.Transact(ctx2, newMessage(defaultAddr, burnMockErrorIdentity, assetID3)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}
