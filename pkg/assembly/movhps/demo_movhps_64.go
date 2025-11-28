package movhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movhps_64.s
var assemblyMovhps64 string

//go:embed stub_movhps_64.go
var stubMovhps64 string

type MOVHPS64 struct {
}

func (v *MOVHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVHPS64) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVHPS64) Name() string {
	return "MOVHPS XMM (2X 64 bit)"
}

func (v *MOVHPS64) Description() string {
	return "Move two floats into the high 64 bits of XMM; low 64 supplied separately."
}

func (v *MOVHPS64) Stub() string {
	return stubMovhps64
}

func (v *MOVHPS64) Assembly() string {
	return assemblyMovhps64
}

func (v *MOVHPS64) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movhps64(&lower, &upper, &ret)

	log.Printf("MOVHPS64 input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
