package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_all_three.s
var assemblyVpermilps256All_three string

//go:embed stub_vpermilps_256_all_three.go
var stubVpermilps256All_three string

type VPERMILPS256ALL_THREE struct {
}

func (v *VPERMILPS256ALL_THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERMILPS256ALL_THREE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256ALL_THREE) Name() string {
	return "VPERMILPS (256 bit) all_three"
}

func (v *VPERMILPS256ALL_THREE) Description() string {
	return "Permute with imm8=0xFF per 128-bit lane: broadcast lane3 element."
}

func (v *VPERMILPS256ALL_THREE) Stub() string {
	return stubVpermilps256All_three
}

func (v *VPERMILPS256ALL_THREE) Assembly() string {
	return assemblyVpermilps256All_three
}

func (v *VPERMILPS256ALL_THREE) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vpermilps256All_three(&vals, &ret)

	log.Printf("VPERMILPS256ALL_THREE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256ALL_THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
