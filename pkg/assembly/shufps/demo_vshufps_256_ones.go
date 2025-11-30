package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_ones.s
var assemblyVshufps256Ones string

//go:embed stub_vshufps_256_ones.go
var stubVshufps256Ones string

type VSHUFPS256ONES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256ONES() *VSHUFPS256ONES {
	return &VSHUFPS256ONES{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256ONES) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256ONES) Name() string {
	return "VSHUFPS (256 bit) ones"
}

func (v *VSHUFPS256ONES) Description() string {
	return "VSHUFPS imm8=0x55 per 128-bit lane: dst = [a1,a1,b1,b1 | a5,a5,b5,b5]"
}

func (v *VSHUFPS256ONES) Stub() string {
	return stubVshufps256Ones
}

func (v *VSHUFPS256ONES) Assembly() string {
	return assemblyVshufps256Ones
}

func (v *VSHUFPS256ONES) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Ones(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256ONES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VSHUFPS256ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
