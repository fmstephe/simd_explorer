package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_higha_higha.s
var assemblyVperm2f128128Higha_higha string

//go:embed stub_vperm2f128_128_higha_higha.go
var stubVperm2f128128Higha_higha string

type VPERM2F128128HIGHA_HIGHA struct {
}

func (v *VPERM2F128128HIGHA_HIGHA) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128HIGHA_HIGHA) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128HIGHA_HIGHA) Name() string {
	return "VPERM2F128 (128 bit) higha_higha"
}

func (v *VPERM2F128128HIGHA_HIGHA) Description() string {
	return "dst.low = A.high, dst.high = A.high (per 128-bit lane)"
}

func (v *VPERM2F128128HIGHA_HIGHA) Stub() string {
	return stubVperm2f128128Higha_higha
}

func (v *VPERM2F128128HIGHA_HIGHA) Assembly() string {
	return assemblyVperm2f128128Higha_higha
}

func (v *VPERM2F128128HIGHA_HIGHA) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Higha_higha(&a, &b, &ret)

	log.Printf("VPERM2F128128HIGHA_HIGHA A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128HIGHA_HIGHA) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
