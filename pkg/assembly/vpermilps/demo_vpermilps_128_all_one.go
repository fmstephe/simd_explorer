package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_all_one.s
var assemblyVpermilps128All_one string

//go:embed stub_vpermilps_128_all_one.go
var stubVpermilps128All_one string

type VPERMILPS128ALL_ONE struct {
}

func (v *VPERMILPS128ALL_ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128ALL_ONE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128ALL_ONE) Name() string {
	return "VPERMILPS (128 bit) all_one"
}

func (v *VPERMILPS128ALL_ONE) Description() string {
	return "Permute with imm8=0x55: broadcast a1 to all lanes."
}

func (v *VPERMILPS128ALL_ONE) Stub() string {
	return stubVpermilps128All_one
}

func (v *VPERMILPS128ALL_ONE) Assembly() string {
	return assemblyVpermilps128All_one
}

func (v *VPERMILPS128ALL_ONE) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128All_one(&vals, &ret)

	log.Printf("VPERMILPS128ALL_ONE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128ALL_ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
