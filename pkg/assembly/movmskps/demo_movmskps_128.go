package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movmskps_128.s
var assemblyMovmskps128 string

//go:embed stub_movmskps_128.go
var stubMovmskps128 string

type MOVMSKPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewMOVMSKPS128() *MOVMSKPS128 {
	return &MOVMSKPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *MOVMSKPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *MOVMSKPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MOVMSKPS128) Name() string {
	return "MOVMSKPS (128 bit) "
}

func (v *MOVMSKPS128) Description() string {
	return "Extract sign bits of packed single-precision elements in XMM into a 4-bit integer mask."
}

func (v *MOVMSKPS128) Stub() string {
	return stubMovmskps128
}

func (v *MOVMSKPS128) Assembly() string {
	return assemblyMovmskps128
}

func (v *MOVMSKPS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]byte{}

	movmskps128(&vals, &ret)

	log.Printf("MOVMSKPS128 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *MOVMSKPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
