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
}

func (v *VMAXSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMAXSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMAXSS128) Name() string {
	return "VMAXSS (128 bit) "
}

func (v *VMAXSS128) Description() string {
	return "TODO"
}

func (v *VMAXSS128) Stub() string {
	return stubVmaxss128
}

func (v *VMAXSS128) Assembly() string {
	return assemblyVmaxss128
}

func (v *VMAXSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmaxss128(&floats1, &floats2, &ret)

	log.Printf("VMAXSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMAXSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
