// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/identities/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/lists"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type identityID struct {
	ids.ClassificationID
	Hash ids.ID
}

var _ ids.IdentityID = (*identityID)(nil)
var _ helpers.Key = (*identityID)(nil)

func (identityID identityID) Bytes() []byte {
	return append(
		identityID.ClassificationID.Bytes(),
		identityID.Hash.Bytes()...,
	)
}
func (identityID identityID) String() string {
	return stringUtilities.JoinIDStrings(identityID.ClassificationID.String(), identityID.Hash.String())
}
func (identityID identityID) Compare(listable traits.Listable) int {
	return bytes.Compare(identityID.Bytes(), identityIDFromInterface(listable).Bytes())
}
func (identityID identityID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(identityID.Bytes())
}
func (identityID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, identityID{})
}
func (identityID identityID) IsPartial() bool {
	return len(identityID.Hash.Bytes()) == 0
}
func (identityID identityID) Equals(key helpers.Key) bool {
	return identityID.Compare(identityIDFromInterface(key)) == 0
}

// TODO Pass Classification & then get Classification ID
func NewIdentityID(classificationID ids.ClassificationID, immutableProperties lists.PropertyList) ids.IdentityID {
	return identityID{
		ClassificationID: classificationID,
		Hash:             baseQualified.immutables{PropertyList: immutableProperties}.GenerateHashID(),
	}
}
