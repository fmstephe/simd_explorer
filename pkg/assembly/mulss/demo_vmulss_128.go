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
}

func (v *VMULSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMULSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
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

func (v *VMULSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmulss128(&floats1, &floats2, &ret)

	log.Printf("VMULSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMULSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
