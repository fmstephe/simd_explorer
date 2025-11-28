package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_128.s
var assemblyVpbroadcastb128 string

//go:embed stub_vpbroadcastb_128.go
var stubVpbroadcastb128 string

type VPBROADCASTB128 struct {
}

func (v *VPBROADCASTB128) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(8, 8, 16)}
}

func (v *VPBROADCASTB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 64, 16)
}

func (v *VPBROADCASTB128) Name() string {
	return "VPBROADCASTB XMM (128 bit)"
}

func (v *VPBROADCASTB128) Description() string {
	return "Broadcast an 8-bit value to all 16 byte elements in XMM."
}

func (v *VPBROADCASTB128) Stub() string {
	return stubVpbroadcastb128
}

func (v *VPBROADCASTB128) Assembly() string {
	return assemblyVpbroadcastb128
}

func (v *VPBROADCASTB128) Run(inputs [][]byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128(inputs[0][0], &ret)
	return ret[:]
}

func (v *VPBROADCASTB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
