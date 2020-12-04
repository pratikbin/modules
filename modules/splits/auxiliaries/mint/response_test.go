package mint

import (
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Mint_Response(t *testing.T) {

	testAuxiliaryResponse := newAuxiliaryResponse(nil)
	require.Equal(t, auxiliaryResponse{Success: true, Error: nil}, testAuxiliaryResponse)
	require.Equal(t, true, testAuxiliaryResponse.IsSuccessful())
	require.Equal(t, nil, testAuxiliaryResponse.GetError())

	testAuxiliaryResponse2 := newAuxiliaryResponse(errors.IncorrectFormat)
	require.Equal(t, auxiliaryResponse{Success: false, Error: errors.IncorrectFormat}, testAuxiliaryResponse2)
	require.Equal(t, false, testAuxiliaryResponse2.IsSuccessful())
	require.Equal(t, errors.IncorrectFormat, testAuxiliaryResponse2.GetError())
}
