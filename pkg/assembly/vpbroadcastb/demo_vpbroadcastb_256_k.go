package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_256_k.s
var assemblyVpbroadcastb256K string

//go:embed stub_vpbroadcastb_256_k.go
var stubVpbroadcastb256K string

type VPBROADCASTB256K struct {
}

func (v *VPBROADCASTB256K) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(8, 8, 16),
		number.NewUintParameter(64, 64, 16),
	}
}

func (v *VPBROADCASTB256K) Output() *number.Parameter {
	return number.NewUintParameter(256, 64, 16)
}

func (v *VPBROADCASTB256K) Name() string {
	return "VPBROADCASTB YMM (256K bit) K"
}

func (v *VPBROADCASTB256K) Description() string {
	return "AVX-512 form: broadcast an 8-bit value to YMM byte elements; lanes written are selected by writemask k."
}

func (v *VPBROADCASTB256K) Stub() string {
	return stubVpbroadcastb256K
}

func (v *VPBROADCASTB256K) Assembly() string {
	return assemblyVpbroadcastb256K
}

func (v *VPBROADCASTB256K) Run(inputs [][]byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb256K(inputs[0][0], number.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *VPBROADCASTB256K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
