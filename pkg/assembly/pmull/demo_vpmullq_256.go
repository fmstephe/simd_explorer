package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmullq_256.s
var assemblyVpmullq256 string

//go:embed stub_vpmullq_256.go
var stubVpmullq256 string

type VPMULLQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLQ256() *VPMULLQ256 {
	return &VPMULLQ256{
		vals1: number.NewNamedUintParameter("vals1", 256, 64, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPMULLQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLQ256) Name() string {
	return "VPMULLQ (256 bit) "
}

func (v *VPMULLQ256) Description() string {
	return "Multiply packed u64 quadwords, keeping the low 64 bits per lane."
}

func (v *VPMULLQ256) Stub() string {
	return stubVpmullq256
}

func (v *VPMULLQ256) Assembly() string {
	return assemblyVpmullq256
}

func (v *VPMULLQ256) Run() {
	vals1 := [4]uint64{}
	copy(vals1[:], number.ToUint64Slice(v.vals1.FlatData()))
	vals2 := [4]uint64{}
	copy(vals2[:], number.ToUint64Slice(v.vals2.FlatData()))

	ret := [4]uint64{}

	vpmullq256(&vals1, &vals2, &ret)

	log.Printf("VPMULLQ256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
