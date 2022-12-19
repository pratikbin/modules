// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/ids"
)

//
//type key struct {
//	ids.ID
//}
//
//var _ helpers.Key = (*key)(nil)
//
//func (key key) RegisterInterfaces(registry types.InterfaceRegistry) {
//	//registry.RegisterInterface("KeyName", (*helpers.Key)(nil))
//}
//
//func (key key) GenerateStoreKeyBytes() []byte {
//	return module.StoreKeyPrefix.GenerateStoreKey(key.Bytes())
//}
//func (key) RegisterCodec(codec *codec.LegacyAmino) {
//	schema.RegisterModuleConcrete(codec, key{})
//}
//func (key key) IsPartial() bool {
//	return key.ID.(*base.ID).GetDataID().HashId == nil
//}
//func (key key) Equals(compareKey helpers.Key) bool {
//	if CompareKey, err := keyFromInterface(compareKey); err != nil {
//		return false
//	} else {
//		return key.ID.Compare(CompareKey.ID) == 0
//	}
//}
//func keyFromInterface(i interface{}) (key, error) {
//	switch value := i.(type) {
//	case key:
//		return value, nil
//	default:
//		return key{}, constants.MetaDataError
//	}
//}

func NewKey(dataID ids.ID) helpers.Key {
	return baseHelpers.NewKey(dataID)
}

func Prototype() helpers.Key { return baseHelpers.PrototypeMetasKey() }
