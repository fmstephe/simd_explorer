package mulps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmulps_128.s
var assemblyVmulps128 string

//go:embed stub_vmulps_128.go
var stubVmulps128 string

type VMULPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMULPS128() *VMULPS128 {
	return &VMULPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMULPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMULPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMULPS128) Name() string {
	return "VMULPS (128 bit) "
}

func (v *VMULPS128) Description() string {
	return "AVX form: multiply packed single-precision floats in XMM, lane-wise."
}

func (v *VMULPS128) Stub() string {
	return stubVmulps128
}

func (v *VMULPS128) Assembly() string {
	return assemblyVmulps128
}

func (v *VMULPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vmulps128(&vals1, &vals2, &ret)

	log.Printf("VMULPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMULPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
