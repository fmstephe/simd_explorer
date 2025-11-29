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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS128IDENTITY() *VPERMILPS128IDENTITY {
	return &VPERMILPS128IDENTITY{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals}
}

func (v *VPERMILPS128IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128IDENTITY) Name() string {
	return "VPERMILPS (128 bit) identity"
}

func (v *VPERMILPS128IDENTITY) Description() string {
	return "Permute single-precision floats with imm8=0xE4 per 128-bit lane: identity."
}

func (v *VPERMILPS128IDENTITY) Stub() string {
	return stubVpermilps128Identity
}

func (v *VPERMILPS128IDENTITY) Assembly() string {
	return assemblyVpermilps128Identity
}

func (v *VPERMILPS128IDENTITY) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float32{}

	vpermilps128Identity(&vals, &ret)

	log.Printf("VPERMILPS128IDENTITY vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPS128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
