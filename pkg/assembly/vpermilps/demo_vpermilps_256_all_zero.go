package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_all_zero.s
var assemblyVpermilps256All_zero string

//go:embed stub_vpermilps_256_all_zero.go
var stubVpermilps256All_zero string

type VPERMILPS256ALL_ZERO struct {
}

func (v *VPERMILPS256ALL_ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERMILPS256ALL_ZERO) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256ALL_ZERO) Name() string {
	return "VPERMILPS (256 bit) all_zero"
}

func (v *VPERMILPS256ALL_ZERO) Description() string {
	return "Permute with imm8=0x00 per 128-bit lane: broadcast lane0 element."
}

func (v *VPERMILPS256ALL_ZERO) Stub() string {
	return stubVpermilps256All_zero
}

func (v *VPERMILPS256ALL_ZERO) Assembly() string {
	return assemblyVpermilps256All_zero
}

func (v *VPERMILPS256ALL_ZERO) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vpermilps256All_zero(&vals, &ret)

	log.Printf("VPERMILPS256ALL_ZERO vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256ALL_ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
