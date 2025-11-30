package comiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcomiss_128.s
var assemblyVcomiss128 string

//go:embed stub_vcomiss_128.go
var stubVcomiss128 string

type VCOMISS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCOMISS128() *VCOMISS128 {
	return &VCOMISS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 32, 32, 2),
	}
}

func (v *VCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCOMISS128) Output() *number.Parameter {
	return v.ret
}

func (v *VCOMISS128) Name() string {
	return "VCOMISS (128 bit)"
}

func (v *VCOMISS128) Description() string {
	return "AVX form: ordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF)."
}

func (v *VCOMISS128) Stub() string {
	return stubVcomiss128
}

func (v *VCOMISS128) Assembly() string {
	return assemblyVcomiss128
}

func (v *VCOMISS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	flags := vcomiss128(&vals1, &vals2)

	log.Printf("VCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		vals1, vals2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	out := []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
	v.ret.SetData(out)

}

func (v *VCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
