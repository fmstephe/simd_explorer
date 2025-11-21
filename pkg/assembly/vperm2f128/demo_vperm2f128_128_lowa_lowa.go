package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowa_lowa.s
var assemblyVperm2f128128Lowa_lowa string

//go:embed stub_vperm2f128_128_lowa_lowa.go
var stubVperm2f128128Lowa_lowa string

type VPERM2F128128LOWA_LOWA struct {
}

func (v *VPERM2F128128LOWA_LOWA) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128LOWA_LOWA) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128LOWA_LOWA) Name() string {
	return "VPERM2F128 (128 bit) lowa_lowa"
}

func (v *VPERM2F128128LOWA_LOWA) Description() string {
	return "dst.low = A.low, dst.high = A.low (per 128-bit lane)"
}

func (v *VPERM2F128128LOWA_LOWA) Stub() string {
	return stubVperm2f128128Lowa_lowa
}

func (v *VPERM2F128128LOWA_LOWA) Assembly() string {
	return assemblyVperm2f128128Lowa_lowa
}

func (v *VPERM2F128128LOWA_LOWA) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Lowa_lowa(&a, &b, &ret)

	log.Printf("VPERM2F128128LOWA_LOWA A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128LOWA_LOWA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
