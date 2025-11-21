package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_256_all_zero.s
var assemblyVpermilpd256All_zero string

//go:embed stub_vpermilpd_256_all_zero.go
var stubVpermilpd256All_zero string

type VPERMILPD256ALL_ZERO struct {
}

func (v *VPERMILPD256ALL_ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),
	}
}

func (v *VPERMILPD256ALL_ZERO) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64)
}

func (v *VPERMILPD256ALL_ZERO) Name() string {
	return "VPERMILPD (256 bit) all_zero"
}

func (v *VPERMILPD256ALL_ZERO) Description() string {
	return "Permute with imm8=0x00 per 128-bit lane: broadcast lane0 element."
}

func (v *VPERMILPD256ALL_ZERO) Stub() string {
	return stubVpermilpd256All_zero
}

func (v *VPERMILPD256ALL_ZERO) Assembly() string {
	return assemblyVpermilpd256All_zero
}

func (v *VPERMILPD256ALL_ZERO) Run(inputs [][]byte) (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(inputs[0]))

	ret := [4]float64{}

	vpermilpd256All_zero(&vals, &ret)

	log.Printf("VPERMILPD256ALL_ZERO vals %v ret %v", vals, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VPERMILPD256ALL_ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
