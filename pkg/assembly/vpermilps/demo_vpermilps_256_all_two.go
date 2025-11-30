package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_all_two.s
var assemblyVpermilps256All_two string

//go:embed stub_vpermilps_256_all_two.go
var stubVpermilps256All_two string

type VPERMILPS256ALL_TWO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS256ALL_TWO() *VPERMILPS256ALL_TWO {
	return &VPERMILPS256ALL_TWO{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VPERMILPS256ALL_TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{v.vals}
}

func (v *VPERMILPS256ALL_TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS256ALL_TWO) Name() string {
	return "VPERMILPS (256 bit) all_two"
}

func (v *VPERMILPS256ALL_TWO) Description() string {
	return "Permute single-precision floats with imm8=0xAA per 128-bit lane: all lanes select element 2."
}

func (v *VPERMILPS256ALL_TWO) Stub() string {
	return stubVpermilps256All_two
}

func (v *VPERMILPS256ALL_TWO) Assembly() string {
	return assemblyVpermilps256All_two
}

func (v *VPERMILPS256ALL_TWO) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [8]float32{}

	vpermilps256All_two(&vals, &ret)

	log.Printf("VPERMILPS256ALL_TWO vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPS256ALL_TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
