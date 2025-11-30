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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS128ALL_ONE() *VPERMILPS128ALL_ONE {
	return &VPERMILPS128ALL_ONE{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128ALL_ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals}
}

func (v *VPERMILPS128ALL_ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128ALL_ONE) Name() string {
	return "VPERMILPS (128 bit) all_one"
}

func (v *VPERMILPS128ALL_ONE) Description() string {
	return "Permute single-precision floats with imm8=0x55 per 128-bit lane: all lanes select element 1."
}

func (v *VPERMILPS128ALL_ONE) Stub() string {
	return stubVpermilps128All_one
}

func (v *VPERMILPS128ALL_ONE) Assembly() string {
	return assemblyVpermilps128All_one
}

func (v *VPERMILPS128ALL_ONE) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float32{}

	vpermilps128All_one(&vals, &ret)

	log.Printf("VPERMILPS128ALL_ONE vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPS128ALL_ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
