package maxps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaxps_128.s
var assemblyVmaxps128 string

//go:embed stub_vmaxps_128.go
var stubVmaxps128 string

type VMAXPS128 struct {
}

func (v *VMAXPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMAXPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMAXPS128) Name() string {
	return "VMAXPS XMM (128 bit)"
}

func (v *VMAXPS128) Description() string {
	return "TODO"
}

func (v *VMAXPS128) Stub() string {
	return stubVmaxps128
}

func (v *VMAXPS128) Assembly() string {
	return assemblyVmaxps128
}

func (v *VMAXPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmaxps128(&floats1, &floats2, &ret)

	log.Printf("VMAXPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMAXPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
