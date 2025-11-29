package maxss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaxss_128.s
var assemblyVmaxss128 string

//go:embed stub_vmaxss_128.go
var stubVmaxss128 string

type VMAXSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMAXSS128() *VMAXSS128 {
	return &VMAXSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMAXSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMAXSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMAXSS128) Name() string {
	return "VMAXSS (128 bit) "
}

func (v *VMAXSS128) Description() string {
	return "AVX form: compute maximum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VMAXSS128) Stub() string {
	return stubVmaxss128
}

func (v *VMAXSS128) Assembly() string {
	return assemblyVmaxss128
}

func (v *VMAXSS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vmaxss128(&vals1, &vals2, &ret)

	log.Printf("VMAXSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMAXSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
