package mulss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmulss_128.s
var assemblyVmulss128 string

//go:embed stub_vmulss_128.go
var stubVmulss128 string

type VMULSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMULSS128() *VMULSS128 {
	return &VMULSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMULSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMULSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMULSS128) Name() string {
	return "VMULSS (128 bit) "
}

func (v *VMULSS128) Description() string {
	return "AVX form: multiply scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VMULSS128) Stub() string {
	return stubVmulss128
}

func (v *VMULSS128) Assembly() string {
	return assemblyVmulss128
}

func (v *VMULSS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vmulss128(&vals1, &vals2, &ret)

	log.Printf("VMULSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMULSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
