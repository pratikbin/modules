// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/utilities/transaction"
)

var (
	testMessage = newMessage(fromAccAddress, testFromID, testFromID, testClassificationID)
)

func Test_messageFromInterface(t *testing.T) {
	type args struct {
		msg types.Msg
	}
	tests := []struct {
		name string
		args args
		want message
	}{
		{"+ve", args{testMessage}, message{fromAccAddress, testFromID, testFromID, testClassificationID}},
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
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, types.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	type fields struct {
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   []types.AccAddress
	}{
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, []types.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	type fields struct {
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	type fields struct {
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	type fields struct {
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	type fields struct {
		From             types.AccAddress
		FromID           ids.IdentityID
		ToID             ids.IdentityID
		ClassificationID ids.ClassificationID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{fromAccAddress, testFromID, testFromID, testClassificationID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := message{
				From:             tt.fields.From,
				FromID:           tt.fields.FromID,
				ToID:             tt.fields.ToID,
				ClassificationID: tt.fields.ClassificationID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	type args struct {
		from             types.AccAddress
		fromID           ids.IdentityID
		toID             ids.IdentityID
		classificationID ids.ClassificationID
	}
	tests := []struct {
		name string
		args args
		want types.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testFromID, testClassificationID}, message{fromAccAddress, testFromID, testFromID, testClassificationID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.toID, tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
