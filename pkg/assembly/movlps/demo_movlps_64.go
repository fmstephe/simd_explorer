package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movlps_64.s
var assemblyMovlps64 string

//go:embed stub_movlps_64.go
var stubMovlps64 string

type MOVLPS64 struct {
}

func (v *MOVLPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVLPS64) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVLPS64) Name() string {
	return "MOVLPS XMM (2X 64 bit)"
}

func (v *MOVLPS64) Description() string {
	return "TODO"
}

func (v *MOVLPS64) Stub() string {
	return stubMovlps64
}

func (v *MOVLPS64) Assembly() string {
	return assemblyMovlps64
}

func (v *MOVLPS64) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movlps64(&lower, &upper, &ret)

	log.Printf("MOVLPS64 input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVLPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
