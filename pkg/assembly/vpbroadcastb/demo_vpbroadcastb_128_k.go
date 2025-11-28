package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_128_k.s
var assemblyVpbroadcastb128K string

//go:embed stub_vpbroadcastb_128_k.go
var stubVpbroadcastb128K string

type VPBROADCASTB128K struct {
}

func (v *VPBROADCASTB128K) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(8, 8, 16),
		number.NewUintParameter(64, 64, 16),
	}
}

func (v *VPBROADCASTB128K) Output() *number.Parameter {
	return number.NewUintParameter(128, 64, 16)
}

func (v *VPBROADCASTB128K) Name() string {
	return "VPBROADCASTB XMM (128K bit) K"
}

func (v *VPBROADCASTB128K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to XMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB128K) Stub() string {
	return stubVpbroadcastb128K
}

func (v *VPBROADCASTB128K) Assembly() string {
	return assemblyVpbroadcastb128K
}

func (v *VPBROADCASTB128K) Run(inputs [][]byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128K(inputs[0][0], number.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *VPBROADCASTB128K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
