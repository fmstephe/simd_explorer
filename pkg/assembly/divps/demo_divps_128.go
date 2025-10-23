package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_divps_128.s
var assemblyDivps128 string

//go:embed stub_divps_128.go
var stubDivps128 string

type DIVPS128 struct {
}

func (v *DIVPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *DIVPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *DIVPS128) Name() string {
	return "DIVPS XMM (128 bit)"
}

func (v *DIVPS128) Description() string {
	return "TODO"
}

func (v *DIVPS128) Stub() string {
	return stubDivps128
}

func (v *DIVPS128) Assembly() string {
	return assemblyDivps128
}

func (v *DIVPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	divps128(&floats1, &floats2, &ret)

	log.Printf("DIVPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *DIVPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
