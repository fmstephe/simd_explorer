package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_128_identity.s
var assemblyVpermilpd128Identity string

//go:embed stub_vpermilpd_128_identity.go
var stubVpermilpd128Identity string

type VPERMILPD128IDENTITY struct {
}

func (v *VPERMILPD128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),
	}
}

func (v *VPERMILPD128IDENTITY) Output() *number.Parameter {
	return number.NewFloatParameter(128, 64)
}

func (v *VPERMILPD128IDENTITY) Name() string {
	return "VPERMILPD (128 bit) identity"
}

func (v *VPERMILPD128IDENTITY) Description() string {
	return "Permute double-precision floats with imm8=0xE4: identity per 128-bit lane."
}

func (v *VPERMILPD128IDENTITY) Stub() string {
	return stubVpermilpd128Identity
}

func (v *VPERMILPD128IDENTITY) Assembly() string {
	return assemblyVpermilpd128Identity
}

func (v *VPERMILPD128IDENTITY) Run(inputs [][]byte) (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(inputs[0]))

	ret := [2]float64{}

	vpermilpd128Identity(&vals, &ret)

	log.Printf("VPERMILPD128IDENTITY vals %v ret %v", vals, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VPERMILPD128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
