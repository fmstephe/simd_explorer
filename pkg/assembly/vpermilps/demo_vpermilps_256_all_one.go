package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_all_one.s
var assemblyVpermilps256All_one string

//go:embed stub_vpermilps_256_all_one.go
var stubVpermilps256All_one string

type VPERMILPS256ALL_ONE struct {
}

func (v *VPERMILPS256ALL_ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERMILPS256ALL_ONE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256ALL_ONE) Name() string {
	return "VPERMILPS (256 bit) all_one"
}

func (v *VPERMILPS256ALL_ONE) Description() string {
	return "Permute with imm8=0x55 per 128-bit lane: broadcast lane1 element."
}

func (v *VPERMILPS256ALL_ONE) Stub() string {
	return stubVpermilps256All_one
}

func (v *VPERMILPS256ALL_ONE) Assembly() string {
	return assemblyVpermilps256All_one
}

func (v *VPERMILPS256ALL_ONE) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vpermilps256All_one(&vals, &ret)

	log.Printf("VPERMILPS256ALL_ONE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256ALL_ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
