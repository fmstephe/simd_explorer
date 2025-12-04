package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmullw_256.s
var assemblyVpmullw256 string

//go:embed stub_vpmullw_256.go
var stubVpmullw256 string

type VPMULLW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLW256() *VPMULLW256 {
	return &VPMULLW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPMULLW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLW256) Name() string {
	return "VPMULLW (256 bit) "
}

func (v *VPMULLW256) Description() string {
	return "Multiply packed u16 words, keeping the low 16 bits per lane."
}

func (v *VPMULLW256) Stub() string {
	return stubVpmullw256
}

func (v *VPMULLW256) Assembly() string {
	return assemblyVpmullw256
}

func (v *VPMULLW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpmullw256(&vals1, &vals2, &ret)

	log.Printf("VPMULLW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
