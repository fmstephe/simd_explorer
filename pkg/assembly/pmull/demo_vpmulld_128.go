package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulld_128.s
var assemblyVpmulld128 string

//go:embed stub_vpmulld_128.go
var stubVpmulld128 string

type VPMULLD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLD128() *VPMULLD128 {
	return &VPMULLD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMULLD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLD128) Name() string {
	return "VPMULLD (128 bit) "
}

func (v *VPMULLD128) Description() string {
	return "Multiply packed u32 doublewords, keeping the low 32 bits per lane."
}

func (v *VPMULLD128) Stub() string {
	return stubVpmulld128
}

func (v *VPMULLD128) Assembly() string {
	return assemblyVpmulld128
}

func (v *VPMULLD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpmulld128(&vals1, &vals2, &ret)

	log.Printf("VPMULLD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
