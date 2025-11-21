package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_all_three.s
var assemblyVpermilps128All_three string

//go:embed stub_vpermilps_128_all_three.go
var stubVpermilps128All_three string

type VPERMILPS128ALL_THREE struct {
}

func (v *VPERMILPS128ALL_THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128ALL_THREE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128ALL_THREE) Name() string {
	return "VPERMILPS (128 bit) all_three"
}

func (v *VPERMILPS128ALL_THREE) Description() string {
	return "Permute with imm8=0xFF: broadcast a3 to all lanes."
}

func (v *VPERMILPS128ALL_THREE) Stub() string {
	return stubVpermilps128All_three
}

func (v *VPERMILPS128ALL_THREE) Assembly() string {
	return assemblyVpermilps128All_three
}

func (v *VPERMILPS128ALL_THREE) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128All_three(&vals, &ret)

	log.Printf("VPERMILPS128ALL_THREE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128ALL_THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
