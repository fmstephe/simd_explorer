package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskps_128.s
var assemblyVmovmskps128 string

//go:embed stub_vmovmskps_128.go
var stubVmovmskps128 string

type VMOVMSKPS128 struct {
}

func (v *VMOVMSKPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVMSKPS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VMOVMSKPS128) Name() string {
	return "VMOVMSKPS (128 bit) "
}

func (v *VMOVMSKPS128) Description() string {
	return "TODO"
}

func (v *VMOVMSKPS128) Stub() string {
	return stubVmovmskps128
}

func (v *VMOVMSKPS128) Assembly() string {
	return assemblyVmovmskps128
}

func (v *VMOVMSKPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]byte{}

	vmovmskps128(&floats, &ret)

	log.Printf("VMOVMSKPS128 input %v output %v", floats, ret)

	return ret[:]
}

func (v *VMOVMSKPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
