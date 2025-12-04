package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmullq_128.s
var assemblyVpmullq128 string

//go:embed stub_vpmullq_128.go
var stubVpmullq128 string

type VPMULLQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLQ128() *VPMULLQ128 {
	return &VPMULLQ128{
		vals1: number.NewNamedUintParameter("vals1", 128, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMULLQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLQ128) Name() string {
	return "VPMULLQ (128 bit) "
}

func (v *VPMULLQ128) Description() string {
	return "Multiply packed u64 quadwords, keeping the low 64 bits per lane."
}

func (v *VPMULLQ128) Stub() string {
	return stubVpmullq128
}

func (v *VPMULLQ128) Assembly() string {
	return assemblyVpmullq128
}

func (v *VPMULLQ128) Run() {
	v.Supported()
	vals1 := [2]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [2]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [2]uint64{}

	vpmullq128(&vals1, &vals2, &ret)

	log.Printf("VPMULLQ128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
