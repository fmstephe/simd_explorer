package vperm2f128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vperm2f128_128_lowb_lowb.s
var assemblyVperm2f128128Lowb_lowb string

//go:embed stub_vperm2f128_128_lowb_lowb.go
var stubVperm2f128128Lowb_lowb string

type VPERM2F128128LOWB_LOWB struct {
}

func (v *VPERM2F128128LOWB_LOWB) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERM2F128128LOWB_LOWB) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERM2F128128LOWB_LOWB) Name() string {
	return "VPERM2F128 (128 bit) lowb_lowb"
}

func (v *VPERM2F128128LOWB_LOWB) Description() string {
	return "dst.low = B.low, dst.high = B.low (per 128-bit lane)"
}

func (v *VPERM2F128128LOWB_LOWB) Stub() string {
	return stubVperm2f128128Lowb_lowb
}

func (v *VPERM2F128128LOWB_LOWB) Assembly() string {
	return assemblyVperm2f128128Lowb_lowb
}

func (v *VPERM2F128128LOWB_LOWB) Run(inputs [][]byte) (output []byte) {
	a := [8]float32{}
	copy(a[:], number.ToFloat32Slice(inputs[0]))
	b := [8]float32{}
	copy(b[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vperm2f128128Lowb_lowb(&a, &b, &ret)

	log.Printf("VPERM2F128128LOWB_LOWB A %v B %v output %v", a, b, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERM2F128128LOWB_LOWB) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
