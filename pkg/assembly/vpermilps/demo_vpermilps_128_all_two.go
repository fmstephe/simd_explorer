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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS128ALL_TWO() *VPERMILPS128ALL_TWO {
	return &VPERMILPS128ALL_TWO{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128ALL_TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals}
}

func (v *VPERMILPS128ALL_TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128ALL_TWO) Name() string {
	return "VPERMILPS (128 bit) all_two"
}

func (v *VPERMILPS128ALL_TWO) Description() string {
	return "Permute single-precision floats with imm8=0xAA per 128-bit lane: all lanes select element 2."
}

func (v *VPERMILPS128ALL_TWO) Stub() string {
	return stubVpermilps128All_two
}

func (v *VPERMILPS128ALL_TWO) Assembly() string {
	return assemblyVpermilps128All_two
}

func (v *VPERMILPS128ALL_TWO) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float32{}

	vpermilps128All_two(&vals, &ret)

	log.Printf("VPERMILPS128ALL_TWO vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPS128ALL_TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
