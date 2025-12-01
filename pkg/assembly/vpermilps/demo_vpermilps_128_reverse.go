package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_reverse.s
var assemblyVpermilps128Reverse string

//go:embed stub_vpermilps_128_reverse.go
var stubVpermilps128Reverse string

type VPERMILPS128REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPS128REVERSE() *VPERMILPS128REVERSE {
	return &VPERMILPS128REVERSE{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VPERMILPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPS128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPS128REVERSE) Name() string {
	return "VPERMILPS (128 bit) reverse"
}

func (v *VPERMILPS128REVERSE) Description() string {
	return "Permute single-precision floats with imm8=0x1B per 128-bit lane: reverse order."
}

func (v *VPERMILPS128REVERSE) Stub() string {
	return stubVpermilps128Reverse
}

func (v *VPERMILPS128REVERSE) Assembly() string {
	return assemblyVpermilps128Reverse
}

func (v *VPERMILPS128REVERSE) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float32{}

	vpermilps128Reverse(&vals, &ret)

	log.Printf("VPERMILPS128REVERSE vals %v ret %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
