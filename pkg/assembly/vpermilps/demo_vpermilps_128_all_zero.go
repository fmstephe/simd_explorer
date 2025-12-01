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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS128ALL_ZERO() *VPERMILPS128ALL_ZERO {
	return &VPERMILPS128ALL_ZERO{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128ALL_ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPS128ALL_ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128ALL_ZERO) Name() string {
	return "VPERMILPS (128 bit) all_zero"
}

func (v *VPERMILPS128ALL_ZERO) Description() string {
	return "Permute single-precision floats with imm8=0x00 per 128-bit lane: all lanes select element 0."
}

func (v *VPERMILPS128ALL_ZERO) Stub() string {
	return stubVpermilps128All_zero
}

func (v *VPERMILPS128ALL_ZERO) Assembly() string {
	return assemblyVpermilps128All_zero
}

func (v *VPERMILPS128ALL_ZERO) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float32{}

	vpermilps128All_zero(&vals, &ret)

	log.Printf("VPERMILPS128ALL_ZERO vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPS128ALL_ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
