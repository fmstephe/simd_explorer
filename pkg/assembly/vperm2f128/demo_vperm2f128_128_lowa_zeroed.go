package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowa_zeroed.s
var assemblyVperm2f128128Lowa_zeroed string

//go:embed stub_vperm2f128_128_lowa_zeroed.go
var stubVperm2f128128Lowa_zeroed string

type VPERM2F128128LOWA_ZEROED struct {
}

func (v *VPERM2F128128LOWA_ZEROED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128LOWA_ZEROED) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128LOWA_ZEROED) Name() string {
	return "VPERM2F128 (128 bit) lowa_zeroed"
}

func (v *VPERM2F128128LOWA_ZEROED) Description() string {
	return "dst.low = A.low, dst.high = zero (per 128-bit lane)"
}

func (v *VPERM2F128128LOWA_ZEROED) Stub() string {
	return stubVperm2f128128Lowa_zeroed
}

func (v *VPERM2F128128LOWA_ZEROED) Assembly() string {
	return assemblyVperm2f128128Lowa_zeroed
}

func (v *VPERM2F128128LOWA_ZEROED) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Lowa_zeroed(&a, &b, &ret)

	log.Printf("VPERM2F128128LOWA_ZEROED A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128LOWA_ZEROED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
