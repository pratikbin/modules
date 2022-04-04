// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Prototype(t *testing.T) {
	prototype := Prototype()
	require.Equal(t, baseHelpers.NewParameters(dummy.Parameter).String(), prototype.String())
	require.Equal(t, nil, prototype.Validate())
	require.Equal(t, dummy.Parameter.String(), prototype.Get(base.NewID("dummy")).String())
}
