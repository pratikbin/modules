// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mutate

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/assets/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	immutableMetaProperties = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutables              = baseQualified.NewImmutables(immutableMetaProperties)
	mutableMetaProperties   = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutableProperties       = baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList())))
	mutables                = baseQualified.NewMutables(mutableMetaProperties)
	classificationID        = baseIDs.NewClassificationID(immutables, mutables)
	fromAddress             = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _       = sdkTypes.AccAddressFromBech32(fromAddress)
	fromID                  = baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID             = baseIDs.NewAssetID(classificationID, immutables)
)

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		{"+ve", args{newMessage(fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties)}, newMessage(fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties).(message)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messageFromInterface(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{
		{"+ve", message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSignBytes(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(newMessage(fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties)))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	type fields struct {
		From                  sdkTypes.AccAddress
		FromID                ids.IdentityID
		AssetID               ids.AssetID
		MutableMetaProperties lists.PropertyList
		MutableProperties     lists.PropertyList
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, false},
		{"+ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:                  tt.fields.From,
				FromID:                tt.fields.FromID,
				AssetID:               tt.fields.AssetID,
				MutableMetaProperties: tt.fields.MutableMetaProperties,
				MutableProperties:     tt.fields.MutableProperties,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	type args struct {
		from                  sdkTypes.AccAddress
		fromID                ids.IdentityID
		assetID               ids.AssetID
		mutableMetaProperties lists.PropertyList
		mutableProperties     lists.PropertyList
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve with nil", args{}, message{}},
		{"+ve", args{fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties}, newMessage(fromAccAddress, fromID, testAssetID, mutableMetaProperties, mutableProperties)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.assetID, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
