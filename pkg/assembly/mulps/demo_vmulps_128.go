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
}

func (v *VMULPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMULPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMULPS128) Name() string {
	return "VMULPS (128 bit) "
}

func (v *VMULPS128) Description() string {
	return "TODO"
}

func (v *VMULPS128) Stub() string {
	return stubVmulps128
}

func (v *VMULPS128) Assembly() string {
	return assemblyVmulps128
}

func (v *VMULPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmulps128(&floats1, &floats2, &ret)

	log.Printf("VMULPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMULPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
