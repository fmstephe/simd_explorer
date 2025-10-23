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
}

func (v *VMULPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMULPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMULPS256) Name() string {
	return "VMULPS (256 bit) "
}

func (v *VMULPS256) Description() string {
	return "TODO"
}

func (v *VMULPS256) Stub() string {
	return stubVmulps256
}

func (v *VMULPS256) Assembly() string {
	return assemblyVmulps256
}

func (v *VMULPS256) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vmulps256(&floats1, &floats2, &ret)

	log.Printf("VMULPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMULPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
