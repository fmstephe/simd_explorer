package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_128_reverse.s
var assemblyVpermilpd128Reverse string

//go:embed stub_vpermilpd_128_reverse.go
var stubVpermilpd128Reverse string

type VPERMILPD128REVERSE struct {
}

func (v *VPERMILPD128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),
	}
}

func (v *VPERMILPD128REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64)
}

func (v *VPERMILPD128REVERSE) Name() string {
	return "VPERMILPD (128 bit) reverse"
}

func (v *VPERMILPD128REVERSE) Description() string {
	return "Permute with imm8=0x1B: reverse the two 64-bit elements [a1 a0]."
}

func (v *VPERMILPD128REVERSE) Stub() string {
	return stubVpermilpd128Reverse
}

func (v *VPERMILPD128REVERSE) Assembly() string {
	return assemblyVpermilpd128Reverse
}

func (v *VPERMILPD128REVERSE) Run(inputs [][]byte) (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(inputs[0]))

	ret := [2]float64{}

	vpermilpd128Reverse(&vals, &ret)

	log.Printf("VPERMILPD128REVERSE vals %v ret %v", vals, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VPERMILPD128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
