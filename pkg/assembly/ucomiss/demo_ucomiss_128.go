package ucomiss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_ucomiss_128.s
var assemblyComiss128 string

//go:embed stub_ucomiss_128.go
var stubComiss128 string

type UCOMISS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewUCOMISS128() *UCOMISS128 {
	return &UCOMISS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		// Flags output (ZF, PF, CF) displayed as hex
		ret: number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *UCOMISS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *UCOMISS128) Output() *number.Parameter {
	return v.ret
}

func (v *UCOMISS128) Name() string {
	return "UCOMISS (128 bit)"
}

func (v *UCOMISS128) Description() string {
	return "Unordered compare scalar single-precision (lane 0); writes EFLAGS (ZF, PF, CF). Upper lanes are ignored."
}

func (v *UCOMISS128) Stub() string {
	return stubComiss128
}

func (v *UCOMISS128) Assembly() string {
	return assemblyComiss128
}

func (v *UCOMISS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	flags := ucomiss128(&vals1, &vals2)

	log.Printf("UCOMISS128 input %v %v output flags=0x%08X [ZF=%d PF=%d CF=%d]",
		vals1, vals2, flags, (flags>>16)&0xFF, (flags>>8)&0xFF, flags&0xFF)

	out := []byte{
		byte(flags),
		byte(flags >> 8),
		byte(flags >> 16),
		byte(flags >> 24),
	}
	v.ret.SetData(out)
	return out
}

func (v *UCOMISS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
