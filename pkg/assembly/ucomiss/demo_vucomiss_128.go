package ucomiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vucomiss_128.s
var assemblyVucomiss128 string

//go:embed stub_vucomiss_128.go
var stubVucomiss128 string

type VUCOMISS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUCOMISS128() *VUCOMISS128 {
	return &VUCOMISS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		// Flags output (ZF, PF, CF) displayed as hex
		ret: number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VUCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUCOMISS128) Output() *number.Parameter {
	return v.ret
}

func (v *VUCOMISS128) Name() string {
	return "VUCOMISS (128 bit)"
}

func (v *VUCOMISS128) Description() string {
	return "AVX form of unordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF)."
}

func (v *VUCOMISS128) Stub() string {
	return stubVucomiss128
}

func (v *VUCOMISS128) Assembly() string {
	return assemblyVucomiss128
}

func (v *VUCOMISS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	flags := vucomiss128(&vals1, &vals2)

	log.Printf("VUCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		vals1, vals2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	out := []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
	v.ret.SetData(out)

}

func (v *VUCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
