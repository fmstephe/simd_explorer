package comiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_comiss_128.s
var assemblyComiss128 string

//go:embed stub_comiss_128.go
var stubComiss128 string

type COMISS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCOMISS128() *COMISS128 {
	return &COMISS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 2),
	}
}

func (v *COMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *COMISS128) Output() *number.Parameter {
	return v.ret
}

func (v *COMISS128) Name() string {
	return "COMISS (128 bit)"
}

func (v *COMISS128) Description() string {
	return "Ordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF)."
}

func (v *COMISS128) Stub() string {
	return stubComiss128
}

func (v *COMISS128) Assembly() string {
	return assemblyComiss128
}

func (v *COMISS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	flags := comiss128(&vals1, &vals2)

	log.Printf("COMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		vals1, vals2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	out := []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
	v.ret.SetData(out)

}

func (v *COMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
