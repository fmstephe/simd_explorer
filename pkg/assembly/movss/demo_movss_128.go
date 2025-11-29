package movss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movss_128.s
var assemblyMovss128 string

//go:embed stub_movss_128.go
var stubMovss128 string

type MOVSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewMOVSS128() *MOVSS128 {
	return &MOVSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MOVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *MOVSS128) Output() *number.Parameter {
	return v.ret
}

func (v *MOVSS128) Name() string {
	return "MOVSS XMM (128 bit)"
}

func (v *MOVSS128) Description() string {
	return "Move scalar single-precision (lane 0) between XMM and memory; upper lanes pass through."
}

func (v *MOVSS128) Stub() string {
	return stubMovss128
}

func (v *MOVSS128) Assembly() string {
	return assemblyMovss128
}

func (v *MOVSS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	movss128(&vals, &ret)

	log.Printf("MOVSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MOVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
