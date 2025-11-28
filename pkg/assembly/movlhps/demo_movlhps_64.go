package movlhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movlhps_64.s
var assemblyMovlhps64 string

//go:embed stub_movlhps_64.go
var stubMovlhps64 string

type MOVLHPS64 struct {
}

func (v *MOVLHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVLHPS64) Output() *number.Parameter {
	return number.NewFloatParameter(64, 32)
}

func (v *MOVLHPS64) Name() string {
	return "MOVLHPS (64 bit) "
}

func (v *MOVLHPS64) Description() string {
	return "Move low 64 bits of source into high 64 of destination XMM; low 64 preserved."
}

func (v *MOVLHPS64) Stub() string {
	return stubMovlhps64
}

func (v *MOVLHPS64) Assembly() string {
	return assemblyMovlhps64
}

func (v *MOVLHPS64) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [2]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [2]float32{}

	movlhps64(&floats, &ret)

	log.Printf("MOVLHPS64 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVLHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
