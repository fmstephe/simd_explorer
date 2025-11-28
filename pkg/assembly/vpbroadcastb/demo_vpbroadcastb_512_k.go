package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_512_k.s
var assemblyVpbroadcastb512K string

//go:embed stub_vpbroadcastb_512_k.go
var stubVpbroadcastb512K string

type VPBROADCASTB512K struct {
}

func (v *VPBROADCASTB512K) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(8, 8, 16),
		number.NewUintParameter(64, 64, 16),
	}
}

func (v *VPBROADCASTB512K) Output() *number.Parameter {
	return number.NewUintParameter(512, 64, 16)
}

func (v *VPBROADCASTB512K) Name() string {
	return "VPBROADCASTB ZMM (512K bit) K"
}

func (v *VPBROADCASTB512K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to ZMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB512K) Stub() string {
	return stubVpbroadcastb512K
}

func (v *VPBROADCASTB512K) Assembly() string {
	return assemblyVpbroadcastb512K
}

func (v *VPBROADCASTB512K) Run(inputs [][]byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512K(inputs[0][0], number.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *VPBROADCASTB512K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
