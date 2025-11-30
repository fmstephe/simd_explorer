package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_divps_128.s
var assemblyDivps128 string

//go:embed stub_divps_128.go
var stubDivps128 string

type DIVPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewDIVPS128() *DIVPS128 {
	return &DIVPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *DIVPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *DIVPS128) Output() *number.Parameter {
	return v.ret
}

func (v *DIVPS128) Name() string {
	return "DIVPS XMM (128 bit)"
}

func (v *DIVPS128) Description() string {
	return "Divide packed single-precision floats in XMM, lane-wise (dest / src)."
}

func (v *DIVPS128) Stub() string {
	return stubDivps128
}

func (v *DIVPS128) Assembly() string {
	return assemblyDivps128
}

func (v *DIVPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	divps128(&vals1, &vals2, &ret)

	log.Printf("DIVPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *DIVPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
