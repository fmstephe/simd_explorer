package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_identity.s
var assemblyVpermilps256Identity string

//go:embed stub_vpermilps_256_identity.go
var stubVpermilps256Identity string

type VPERMILPS256IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS256IDENTITY() *VPERMILPS256IDENTITY {
	return &VPERMILPS256IDENTITY{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERMILPS256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPS256IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS256IDENTITY) Name() string {
	return "VPERMILPS (256 bit) identity"
}

func (v *VPERMILPS256IDENTITY) Description() string {
	return "Permute single-precision floats with imm8=0xE4 per 128-bit lane: identity."
}

func (v *VPERMILPS256IDENTITY) Stub() string {
	return stubVpermilps256Identity
}

func (v *VPERMILPS256IDENTITY) Assembly() string {
	return assemblyVpermilps256Identity
}

func (v *VPERMILPS256IDENTITY) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [8]float32{}

	vpermilps256Identity(&vals, &ret)

	log.Printf("VPERMILPS256IDENTITY vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPS256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
