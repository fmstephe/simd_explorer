package mulss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_mulss_128.s
var assemblyMulss128 string

//go:embed stub_mulss_128.go
var stubMulss128 string

type MULSS128 struct {
}

func (v *MULSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MULSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MULSS128) Name() string {
	return "MULSS (128 bit) "
}

func (v *MULSS128) Description() string {
	return "TODO"
}

func (v *MULSS128) Stub() string {
	return stubMulss128
}

func (v *MULSS128) Assembly() string {
	return assemblyMulss128
}

func (v *MULSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	mulss128(&floats1, &floats2, &ret)

	log.Printf("MULSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MULSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}