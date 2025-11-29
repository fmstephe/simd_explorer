package maxps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaxps_256.s
var assemblyVmaxps256 string

//go:embed stub_vmaxps_256.go
var stubVmaxps256 string

type VMAXPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMAXPS256() *VMAXPS256 {
	return &VMAXPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMAXPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMAXPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMAXPS256) Name() string {
	return "VMAXPS YMM (256 bit)"
}

func (v *VMAXPS256) Description() string {
	return "AVX form: compute element-wise maximum of packed single-precision floats in YMM."
}

func (v *VMAXPS256) Stub() string {
	return stubVmaxps256
}

func (v *VMAXPS256) Assembly() string {
	return assemblyVmaxps256
}

func (v *VMAXPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vmaxps256(&vals1, &vals2, &ret)

	log.Printf("VMAXPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMAXPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
