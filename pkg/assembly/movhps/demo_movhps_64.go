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
	lower *number.Parameter
	upper *number.Parameter
	ret   *number.Parameter
}

func NewMOVHPS64() *MOVHPS64 {
	return &MOVHPS64{
		lower: number.NewNamedFloatParameter("lower", 64, 32),
		upper: number.NewNamedFloatParameter("upper", 64, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MOVHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.lower,
		v.upper,
	}
}

func (v *MOVHPS64) Output() *number.Parameter {
	return v.ret
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

func (v *MOVHPS64) Run() (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(v.lower.FlatData()))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(v.upper.FlatData()))

	ret := [4]float32{}

	movhps64(&lower, &upper, &ret)

	log.Printf("MOVHPS64 input lower %v upper %v output %v", lower, upper, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MOVHPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
