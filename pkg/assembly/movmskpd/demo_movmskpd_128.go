package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movmskpd_128.s
var assemblyMovmskpd128 string

//go:embed stub_movmskpd_128.go
var stubMovmskpd128 string

type MOVMSKPD128 struct {
}

func (v *MOVMSKPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),
	}
}

func (v *MOVMSKPD128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *MOVMSKPD128) Name() string {
	return "MOVMSKPD (128 bit) "
}

func (v *MOVMSKPD128) Description() string {
	return "TODO"
}

func (v *MOVMSKPD128) Stub() string {
	return stubMovmskpd128
}

func (v *MOVMSKPD128) Assembly() string {
	return assemblyMovmskpd128
}

func (v *MOVMSKPD128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [2]float64{}
	copy(floats[:], number.ToFloat64Slice(inputs[0]))

	ret := [4]byte{}

	movmskpd128(&floats, &ret)

	log.Printf("MOVMSKPD128 input %v output %v", floats, ret)

	return ret[:]
}

func (v *MOVMSKPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
