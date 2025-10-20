package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskps_256.s
var assemblyVmovmskps256 string

//go:embed stub_vmovmskps_256.go
var stubVmovmskps256 string

type VMOVMSKPS256 struct {
}

func (v *VMOVMSKPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMOVMSKPS256) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VMOVMSKPS256) Name() string {
	return "VMOVMSKPS (256 bit) "
}

func (v *VMOVMSKPS256) Description() string {
	return "TODO"
}

func (v *VMOVMSKPS256) Stub() string {
	return stubVmovmskps256
}

func (v *VMOVMSKPS256) Assembly() string {
	return assemblyVmovmskps256
}

func (v *VMOVMSKPS256) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]byte{}

	vmovmskps256(&floats, &ret)

	log.Printf("VMOVMSKPS256 input %v output %v", floats, ret)

	return ret[:]
}

func (v *VMOVMSKPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
