package mulps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_mulps_128.s
var assemblyMulps128 string

//go:embed stub_mulps_128.go
var stubMulps128 string

type MULPS128 struct {
}

func (v *MULPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MULPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MULPS128) Name() string {
	return "MULPS (128 bit) "
}

func (v *MULPS128) Description() string {
	return "TODO"
}

func (v *MULPS128) Stub() string {
	return stubMulps128
}

func (v *MULPS128) Assembly() string {
	return assemblyMulps128
}

func (v *MULPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	mulps128(&floats1, &floats2, &ret)

	log.Printf("MULPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MULPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
