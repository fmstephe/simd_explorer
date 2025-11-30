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
	lower *number.Parameter
	upper *number.Parameter
	ret   *number.Parameter
}

func NewMOVLPS64() *MOVLPS64 {
	return &MOVLPS64{
		lower: number.NewNamedFloatParameter("lower", 64, 32),
		upper: number.NewNamedFloatParameter("upper", 64, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MOVLPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.lower,
		v.upper,
	}
}

func (v *MOVLPS64) Output() *number.Parameter {
	return v.ret
}

func (v *MOVLPS64) Name() string {
	return "MOVLPS XMM (2X 64 bit)"
}

func (v *MOVLPS64) Description() string {
	return "Move two floats into the low 64 bits of XMM; high 64 supplied separately."
}

func (v *MOVLPS64) Stub() string {
	return stubMovlps64
}

func (v *MOVLPS64) Assembly() string {
	return assemblyMovlps64
}

func (v *MOVLPS64) Run() {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(v.lower.FlatData()))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(v.upper.FlatData()))

	ret := [4]float32{}

	movlps64(&lower, &upper, &ret)

	log.Printf("MOVLPS64 input lower %v upper %v output %v", lower, upper, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *MOVLPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
