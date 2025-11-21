package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_all_zero.s
var assemblyVpermilps128All_zero string

//go:embed stub_vpermilps_128_all_zero.go
var stubVpermilps128All_zero string

type VPERMILPS128ALL_ZERO struct {
}

func (v *VPERMILPS128ALL_ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128ALL_ZERO) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128ALL_ZERO) Name() string {
	return "VPERMILPS (128 bit) all_zero"
}

func (v *VPERMILPS128ALL_ZERO) Description() string {
	return "Permute with imm8=0x00: broadcast a0 to all lanes."
}

func (v *VPERMILPS128ALL_ZERO) Stub() string {
	return stubVpermilps128All_zero
}

func (v *VPERMILPS128ALL_ZERO) Assembly() string {
	return assemblyVpermilps128All_zero
}

func (v *VPERMILPS128ALL_ZERO) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128All_zero(&vals, &ret)

	log.Printf("VPERMILPS128ALL_ZERO vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128ALL_ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
