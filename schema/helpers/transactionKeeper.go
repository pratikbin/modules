// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type TransactionKeeper interface {
	Transact(sdkTypes.Context, Message) TransactionResponse
	Keeper
}
