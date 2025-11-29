package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivps_128.s
var assemblyVdivps128 string

//go:embed stub_vdivps_128.go
var stubVdivps128 string

type VDIVPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDIVPS128() *VDIVPS128 {
	return &VDIVPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDIVPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDIVPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VDIVPS128) Name() string {
	return "VDIVPS XMM (128 bit)"
}

func (v *VDIVPS128) Description() string {
	return "AVX form: divide packed single-precision floats in XMM, lane-wise (dest / src)."
}

func (v *VDIVPS128) Stub() string {
	return stubVdivps128
}

func (v *VDIVPS128) Assembly() string {
	return assemblyVdivps128
}

func (v *VDIVPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdivps128(&vals1, &vals2, &ret)

	log.Printf("VDIVPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VDIVPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
