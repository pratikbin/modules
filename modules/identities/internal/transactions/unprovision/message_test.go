// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func CreateTestInputForMessage(t *testing.T) (ids.IdentityID, sdkTypes.AccAddress, sdkTypes.AccAddress, sdkTypes.Msg) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentityID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	toAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	toAccAddress, err := sdkTypes.AccAddressFromBech32(toAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, toAccAddress, testIdentityID)

	return testIdentityID, fromAccAddress, toAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, testMessage := CreateTestInputForMessage(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		{"+ve", args{testMessage}, message{fromAccAddress, toAccAddress, testIdentityID}},
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
	testIdentityID, fromAccAddress, toAccAddress, testMessage := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type fields struct {
		From sdkTypes.AccAddress
		To   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress, toAccAddress, testIdentityID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:       tt.fields.From,
				To:         tt.fields.To,
				IdentityID: tt.fields.IdentityID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testIdentityID, fromAccAddress, toAccAddress, _ := CreateTestInputForMessage(t)

	type args struct {
		from sdkTypes.AccAddress
		to   sdkTypes.AccAddress
		ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, toAccAddress, testIdentityID}, message{fromAccAddress, toAccAddress, testIdentityID}},
		{"-ve", args{}, message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.to, tt.args.IdentityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
