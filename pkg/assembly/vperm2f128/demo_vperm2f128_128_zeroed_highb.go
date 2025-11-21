package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_zeroed_highb.s
var assemblyVperm2f128128Zeroed_highb string

//go:embed stub_vperm2f128_128_zeroed_highb.go
var stubVperm2f128128Zeroed_highb string

type VPERM2F128128ZEROED_HIGHB struct {
}

func (v *VPERM2F128128ZEROED_HIGHB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128ZEROED_HIGHB) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128ZEROED_HIGHB) Name() string {
	return "VPERM2F128 (128 bit) zeroed_highb"
}

func (v *VPERM2F128128ZEROED_HIGHB) Description() string {
	return "dst.low = zero, dst.high = B.high (per 128-bit lane)"
}

func (v *VPERM2F128128ZEROED_HIGHB) Stub() string {
	return stubVperm2f128128Zeroed_highb
}

func (v *VPERM2F128128ZEROED_HIGHB) Assembly() string {
	return assemblyVperm2f128128Zeroed_highb
}

func (v *VPERM2F128128ZEROED_HIGHB) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Zeroed_highb(&a, &b, &ret)

	log.Printf("VPERM2F128128ZEROED_HIGHB A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128ZEROED_HIGHB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
