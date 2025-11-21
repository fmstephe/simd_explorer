package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_identity.s
var assemblyVpermilps128Identity string

//go:embed stub_vpermilps_128_identity.go
var stubVpermilps128Identity string

type VPERMILPS128IDENTITY struct {
}

func (v *VPERMILPS128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128IDENTITY) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128IDENTITY) Name() string {
	return "VPERMILPS (128 bit) identity"
}

func (v *VPERMILPS128IDENTITY) Description() string {
	return "Permute single-precision floats with imm8=0xE4: identity per 128-bit lane."
}

func (v *VPERMILPS128IDENTITY) Stub() string {
	return stubVpermilps128Identity
}

func (v *VPERMILPS128IDENTITY) Assembly() string {
	return assemblyVpermilps128Identity
}

func (v *VPERMILPS128IDENTITY) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128Identity(&vals, &ret)

	log.Printf("VPERMILPS128IDENTITY vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
