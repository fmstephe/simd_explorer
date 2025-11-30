package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movups_128.s
var assemblyMovups128 string

//go:embed stub_movups_128.go
var stubMovups128 string

type MOVUPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewMOVUPS128() *MOVUPS128 {
	return &MOVUPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MOVUPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *MOVUPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MOVUPS128) Name() string {
	return "MOVUPS XMM (128 bit)"
}

func (v *MOVUPS128) Description() string {
	return "Unaligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *MOVUPS128) Stub() string {
	return stubMovups128
}

func (v *MOVUPS128) Assembly() string {
	return assemblyMovups128
}

func (v *MOVUPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	movups128(&vals, &ret)

	log.Printf("MOVUPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *MOVUPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
