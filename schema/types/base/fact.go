/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type fact struct {
	HashID     types.ID         `json:"hashID"`
	TypeID     types.ID         `json:"typeID"`
	Signatures types.Signatures `json:"signatures"`
}

var _ types.Fact = (*fact)(nil)

func (fact fact) GetHashID() types.ID             { return fact.HashID }
func (fact fact) GetTypeID() types.ID             { return fact.TypeID }
func (fact fact) GetSignatures() types.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return false
}
func (fact fact) Sign(_ keyring.Keyring) types.Fact {
	// TODO use keyring to sign
	cliContext := context.NewCLIContext()
	sign, _, _ := cliContext.Keybase.Sign(cliContext.FromName, keys.DefaultKeyPass, fact.HashID.Bytes())
	Signature := signature{
		ID:             id{IDString: fact.HashID.String()},
		SignatureBytes: sign,
		ValidityHeight: height{cliContext.Height},
	}
	fact.GetSignatures().Add(Signature)

	return fact
}

func NewFact(data types.Data) types.Fact {
	return fact{
		HashID:     data.GenerateHashID(),
		TypeID:     data.GetTypeID(),
		Signatures: signatures{},
	}
}

func ReadFact(metaFactString string) (types.Fact, error) {
	metaFact, err := ReadMetaFact(metaFactString)
	if err != nil {
		return nil, err
	}

	return metaFact.RemoveData(), nil
}
