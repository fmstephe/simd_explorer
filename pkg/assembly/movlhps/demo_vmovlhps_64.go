package movlhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovlhps_64.s
var assemblyVmovlhps64 string

//go:embed stub_vmovlhps_64.go
var stubVmovlhps64 string

type VMOVLHPS64 struct {
}

func (v *VMOVLHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
	}
}

func (v *VMOVLHPS64) Output() *number.Parameter {
	return number.NewFloatParameter(64, 32)
}

func (v *VMOVLHPS64) Name() string {
	return "VMOVLHPS (64 bit) "
}

func (v *VMOVLHPS64) Description() string {
	return "TODO"
}

func (v *VMOVLHPS64) Stub() string {
	return stubVmovlhps64
}

func (v *VMOVLHPS64) Assembly() string {
	return assemblyVmovlhps64
}

func (v *VMOVLHPS64) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [2]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [2]float32{}

	vmovlhps64(&floats, &ret)

	log.Printf("VMOVLHPS64 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVLHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
