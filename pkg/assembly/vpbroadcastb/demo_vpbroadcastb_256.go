package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_256.s
var assemblyVpbroadcastb256 string

//go:embed stub_vpbroadcastb_256.go
var stubVpbroadcastb256 string

type VPBROADCASTB256 struct {
}

func (v *VPBROADCASTB256) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(8, 8, 16)}
}

func (v *VPBROADCASTB256) Output() *number.Parameter {
	return number.NewUintParameter(256, 64, 16)
}

func (v *VPBROADCASTB256) Name() string {
	return "VPBROADCASTB YMM (256 bit)"
}

func (v *VPBROADCASTB256) Description() string {
	return "Broadcast an 8-bit value to all 32 byte elements in YMM."
}

func (v *VPBROADCASTB256) Stub() string {
	return stubVpbroadcastb256
}

func (v *VPBROADCASTB256) Assembly() string {
	return assemblyVpbroadcastb256
}

func (v *VPBROADCASTB256) Run(inputs [][]byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb256(inputs[0][0], &ret)
	return ret[:]
}

func (v *VPBROADCASTB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
