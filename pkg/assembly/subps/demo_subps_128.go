package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_subps_128.s
var assemblySubps128 string

//go:embed stub_subps_128.go
var stubSubps128 string

type SUBPS128 struct {
}

func (v *SUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SUBPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SUBPS128) Name() string {
	return "SUBPS (128 bit) "
}

func (v *SUBPS128) Description() string {
	return "Subtract packed single-precision floats in XMM, lane-wise (dest - src)."
}

func (v *SUBPS128) Stub() string {
	return stubSubps128
}

func (v *SUBPS128) Assembly() string {
	return assemblySubps128
}

func (v *SUBPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	subps128(&floats1, &floats2, &ret)

	log.Printf("SUBPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
