// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

func createTestData() ids.MaintainerID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	// testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testMaintainerID := baseIDs.NewMaintainerID(testClassificationID, immutables)
	return testMaintainerID
}

func TestNewKey(t *testing.T) {
	type args struct {
		maintainerID ids.MaintainerID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve with nil", args{}, key{}},
		{"+ve", args{createTestData()}, key{createTestData()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.maintainerID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", key{}},
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
		want    key
		wantErr bool
	}{
		{"+ve", args{NewKey(createTestData())}, key{createTestData()}, false},
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
		MaintainerID ids.MaintainerID
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
		{"+ve", fields{createTestData()}, args{NewKey(createTestData())}, true},
		{"-ve", fields{createTestData()}, args{NewKey(baseIDs.PrototypeMaintainerID())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		MaintainerID ids.MaintainerID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{createTestData()}, module.StoreKeyPrefix.GenerateStoreKey(key{createTestData()}.Bytes())},
		{"-ve", fields{baseIDs.PrototypeMaintainerID()}, module.StoreKeyPrefix.GenerateStoreKey(key{baseIDs.PrototypeMaintainerID()}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.GenerateStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		MaintainerID ids.MaintainerID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{createTestData()}, false},
		{"-ve", fields{baseIDs.PrototypeMaintainerID()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := key{
				MaintainerID: tt.fields.MaintainerID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_RegisterCodec(t *testing.T) {
	type fields struct {
		MaintainerID ids.MaintainerID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{createTestData()}, args{codec.NewLegacyAmino()}},
		{"-ve", fields{createTestData()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ke := key{
				MaintainerID: tt.fields.MaintainerID,
			}
			ke.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
