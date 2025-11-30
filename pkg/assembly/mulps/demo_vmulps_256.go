package mulps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmulps_256.s
var assemblyVmulps256 string

//go:embed stub_vmulps_256.go
var stubVmulps256 string

type VMULPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMULPS256() *VMULPS256 {
	return &VMULPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMULPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMULPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMULPS256) Name() string {
	return "VMULPS (256 bit) "
}

func (v *VMULPS256) Description() string {
	return "AVX form: multiply packed single-precision floats in YMM, lane-wise."
}

func (v *VMULPS256) Stub() string {
	return stubVmulps256
}

func (v *VMULPS256) Assembly() string {
	return assemblyVmulps256
}

func (v *VMULPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vmulps256(&vals1, &vals2, &ret)

	log.Printf("VMULPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMULPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
