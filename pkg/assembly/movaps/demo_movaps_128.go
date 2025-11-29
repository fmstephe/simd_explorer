package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movaps_128.s
var assemblyMovaps128 string

//go:embed stub_movaps_128.go
var stubMovaps128 string

type MOVAPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewMOVAPS128() *MOVAPS128 {
	return &MOVAPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MOVAPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *MOVAPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MOVAPS128) Name() string {
	return "MOVAPS XMM (128 bit)"
}

func (v *MOVAPS128) Description() string {
	return "Aligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *MOVAPS128) Stub() string {
	return stubMovaps128
}

func (v *MOVAPS128) Assembly() string {
	return assemblyMovaps128
}

func (v *MOVAPS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	movaps128(&vals, &ret)

	log.Printf("MOVAPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MOVAPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
