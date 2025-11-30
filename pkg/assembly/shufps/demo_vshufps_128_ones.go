package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_ones.s
var assemblyVshufps128Ones string

//go:embed stub_vshufps_128_ones.go
var stubVshufps128Ones string

type VSHUFPS128ONES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS128ONES() *VSHUFPS128ONES {
	return &VSHUFPS128ONES{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSHUFPS128ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS128ONES) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS128ONES) Name() string {
	return "VSHUFPS (128 bit) ones"
}

func (v *VSHUFPS128ONES) Description() string {
	return "VSHUFPS imm8=0x55: dst = [a1,a1,b1,b1]"
}

func (v *VSHUFPS128ONES) Stub() string {
	return stubVshufps128Ones
}

func (v *VSHUFPS128ONES) Assembly() string {
	return assemblyVshufps128Ones
}

func (v *VSHUFPS128ONES) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vshufps128Ones(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS128ONES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VSHUFPS128ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
