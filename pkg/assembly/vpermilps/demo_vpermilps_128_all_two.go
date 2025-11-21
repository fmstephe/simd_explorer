package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_all_two.s
var assemblyVpermilps128All_two string

//go:embed stub_vpermilps_128_all_two.go
var stubVpermilps128All_two string

type VPERMILPS128ALL_TWO struct {
}

func (v *VPERMILPS128ALL_TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128ALL_TWO) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128ALL_TWO) Name() string {
	return "VPERMILPS (128 bit) all_two"
}

func (v *VPERMILPS128ALL_TWO) Description() string {
	return "Permute with imm8=0xAA: broadcast a2 to all lanes."
}

func (v *VPERMILPS128ALL_TWO) Stub() string {
	return stubVpermilps128All_two
}

func (v *VPERMILPS128ALL_TWO) Assembly() string {
	return assemblyVpermilps128All_two
}

func (v *VPERMILPS128ALL_TWO) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128All_two(&vals, &ret)

	log.Printf("VPERMILPS128ALL_TWO vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128ALL_TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
