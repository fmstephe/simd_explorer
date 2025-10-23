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
}

func (v *VMAXPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMAXPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMAXPS256) Name() string {
	return "VMAXPS YMM (256 bit)"
}

func (v *VMAXPS256) Description() string {
	return "TODO"
}

func (v *VMAXPS256) Stub() string {
	return stubVmaxps256
}

func (v *VMAXPS256) Assembly() string {
	return assemblyVmaxps256
}

func (v *VMAXPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vmaxps256(&floats1, &floats2, &ret)

	log.Printf("VMAXPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMAXPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
