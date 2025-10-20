package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovaps_256.s
var assemblyVmovaps256 string

//go:embed stub_vmovaps_256.go
var stubVmovaps256 string

type VMOVAPS256 struct {
}

func (v *VMOVAPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMOVAPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMOVAPS256) Name() string {
	return "VMOVAPS YMM (256 bit)"
}

func (v *VMOVAPS256) Description() string {
	return "TODO"
}

func (v *VMOVAPS256) Stub() string {
	return stubVmovaps256
}

func (v *VMOVAPS256) Assembly() string {
	return assemblyVmovaps256
}

func (v *VMOVAPS256) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vmovaps256(&floats, &ret)

	log.Printf("VMOVAPS256 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVAPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
