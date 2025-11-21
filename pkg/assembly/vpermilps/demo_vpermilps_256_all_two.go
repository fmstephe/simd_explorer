package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_all_two.s
var assemblyVpermilps256All_two string

//go:embed stub_vpermilps_256_all_two.go
var stubVpermilps256All_two string

type VPERMILPS256ALL_TWO struct {
}

func (v *VPERMILPS256ALL_TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERMILPS256ALL_TWO) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256ALL_TWO) Name() string {
	return "VPERMILPS (256 bit) all_two"
}

func (v *VPERMILPS256ALL_TWO) Description() string {
	return "Permute with imm8=0xAA per 128-bit lane: broadcast lane2 element."
}

func (v *VPERMILPS256ALL_TWO) Stub() string {
	return stubVpermilps256All_two
}

func (v *VPERMILPS256ALL_TWO) Assembly() string {
	return assemblyVpermilps256All_two
}

func (v *VPERMILPS256ALL_TWO) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vpermilps256All_two(&vals, &ret)

	log.Printf("VPERMILPS256ALL_TWO vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256ALL_TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
