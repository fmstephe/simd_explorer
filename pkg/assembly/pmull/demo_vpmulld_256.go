package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulld_256.s
var assemblyVpmulld256 string

//go:embed stub_vpmulld_256.go
var stubVpmulld256 string

type VPMULLD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLD256() *VPMULLD256 {
	return &VPMULLD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMULLD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLD256) Name() string {
	return "VPMULLD (256 bit) "
}

func (v *VPMULLD256) Description() string {
	return "Multiply packed u32 doublewords, keeping the low 32 bits per lane."
}

func (v *VPMULLD256) Stub() string {
	return stubVpmulld256
}

func (v *VPMULLD256) Assembly() string {
	return assemblyVpmulld256
}

func (v *VPMULLD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpmulld256(&vals1, &vals2, &ret)

	log.Printf("VPMULLD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
