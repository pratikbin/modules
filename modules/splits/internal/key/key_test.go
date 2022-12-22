// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/splits/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID       = baseIds.NewOwnableID(baseIds.NewStringID("ownerid"))
	splitID             = baseIds.NewSplitID(testOwnerIdentityID, testOwnableID)
)

func TestNewKey(t *testing.T) {
	type args struct {
		splitID *baseIds.SplitID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve", args{splitID.(*baseIds.SplitID)}, &Key{splitID.(*baseIds.SplitID)}},
		{"+ve with nil", args{baseIds.PrototypeSplitID().(*baseIds.SplitID)}, &Key{baseIds.PrototypeSplitID().(*baseIds.SplitID)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.splitID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"+ve", &Key{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyFromInterface(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Key
		wantErr bool
	}{
		{"+ve", args{NewKey(splitID)}, Key{splitID.(*baseIds.SplitID)}, false},
		{"+ve", args{NewKey(baseIds.PrototypeSplitID())}, Key{baseIds.PrototypeSplitID().(*baseIds.SplitID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := keyFromInterface(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("keyFromInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyFromInterface() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_Equals(t *testing.T) {
	type fields struct {
		SplitID *baseIds.SplitID
	}
	type args struct {
		compareKey helpers.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{splitID.(*baseIds.SplitID)}, args{NewKey(splitID)}, true},
		{"+ve", fields{splitID.(*baseIds.SplitID)}, args{NewKey(baseIds.PrototypeSplitID())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				SplitId: tt.fields.SplitID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		SplitID *baseIds.SplitID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{splitID.(*baseIds.SplitID)}, module.StoreKeyPrefix.GenerateStoreKey(splitID.(*baseIds.SplitID).Bytes())},
		{"+ve", fields{baseIds.PrototypeSplitID().(*baseIds.SplitID)}, module.StoreKeyPrefix.GenerateStoreKey(baseIds.PrototypeSplitID().Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				SplitId: tt.fields.SplitID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		SplitID *baseIds.SplitID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{splitID.(*baseIds.SplitID)}, false},
		{"+ve", fields{baseIds.PrototypeSplitID().(*baseIds.SplitID)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				SplitId: tt.fields.SplitID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		SplitID *baseIds.SplitID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{splitID.(*baseIds.SplitID)}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := Key{
				SplitId: tt.fields.SplitID,
			}
			key.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
