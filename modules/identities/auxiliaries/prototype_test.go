package auxiliaries

import (
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"+ve", baseHelpers.NewAuxiliaries(verify.Auxiliary).Get("verify").GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype().Get("verify").GetName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
