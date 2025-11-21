package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_zeroed_zeroed.s
var assemblyVperm2f128128Zeroed_zeroed string

//go:embed stub_vperm2f128_128_zeroed_zeroed.go
var stubVperm2f128128Zeroed_zeroed string

type VPERM2F128128ZEROED_ZEROED struct {
}

func (v *VPERM2F128128ZEROED_ZEROED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128ZEROED_ZEROED) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128ZEROED_ZEROED) Name() string {
	return "VPERM2F128 (128 bit) zeroed_zeroed"
}

func (v *VPERM2F128128ZEROED_ZEROED) Description() string {
	return "dst.low = zero, dst.high = zero (per 128-bit lane)"
}

func (v *VPERM2F128128ZEROED_ZEROED) Stub() string {
	return stubVperm2f128128Zeroed_zeroed
}

func (v *VPERM2F128128ZEROED_ZEROED) Assembly() string {
	return assemblyVperm2f128128Zeroed_zeroed
}

func (v *VPERM2F128128ZEROED_ZEROED) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Zeroed_zeroed(&a, &b, &ret)

	log.Printf("VPERM2F128128ZEROED_ZEROED A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128ZEROED_ZEROED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
