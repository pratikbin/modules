/*
  Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package constants

import "github.com/cosmos/cosmos-sdk/types/errors"

var DeletionNotAllowed = errors.Register(ProjectRoute, 101, "DeletionNotAllowed")
var EntityAlreadyExists = errors.Register(ProjectRoute, 102, "EntityAlreadyExists")
var EntityNotFound = errors.Register(ProjectRoute, 103, "EntityNotFound")
var IncorrectMessage = errors.Register(ProjectRoute, 104, "IncorrectMessage")
var NotAuthorized = errors.Register(ProjectRoute, 106, "NotAuthorized")
