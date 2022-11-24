// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type key struct {
	ids.OrderID
}

var _ helpers.Key = (*key)(nil)

func (key key) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
}
func (key) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(codec, key{})
}
func (key key) IsPartial() bool {
	return len(key.OrderID.Bytes()) == 0
}
func (key key) Equals(compareKey helpers.Key) bool {
	if compareKey, err := keyFromInterface(compareKey); err != nil {
		return false
	} else {
		return key.OrderID.Compare(compareKey.OrderID) == 0
	}
}
func keyFromInterface(i interface{}) (key, error) {
	switch value := i.(type) {
	case key:
		return value, nil
	default:
		return key{}, constants.MetaDataError
	}
}

func NewKey(orderID ids.OrderID) helpers.Key {
	return key{
		OrderID: orderID,
	}
}

func Prototype() helpers.Key {
	return key{}
}
