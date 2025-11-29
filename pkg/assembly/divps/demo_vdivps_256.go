package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivps_256.s
var assemblyVdivps256 string

//go:embed stub_vdivps_256.go
var stubVdivps256 string

type VDIVPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDIVPS256() *VDIVPS256 {
	return &VDIVPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VDIVPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDIVPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VDIVPS256) Name() string {
	return "VDIVPS YMM (256 bit)"
}

func (v *VDIVPS256) Description() string {
	return "AVX form: divide packed single-precision floats in YMM, lane-wise (dest / src)."
}

func (v *VDIVPS256) Stub() string {
	return stubVdivps256
}

func (v *VDIVPS256) Assembly() string {
	return assemblyVdivps256
}

func (v *VDIVPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vdivps256(&vals1, &vals2, &ret)

	log.Printf("VDIVPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VDIVPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
