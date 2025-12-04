package pmull

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmullw_128.s
var assemblyVpmullw128 string

//go:embed stub_vpmullw_128.go
var stubVpmullw128 string

type VPMULLW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULLW128() *VPMULLW128 {
	return &VPMULLW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPMULLW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULLW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULLW128) Name() string {
	return "VPMULLW (128 bit) "
}

func (v *VPMULLW128) Description() string {
	return "Multiply packed u16 words, keeping the low 16 bits per lane."
}

func (v *VPMULLW128) Stub() string {
	return stubVpmullw128
}

func (v *VPMULLW128) Assembly() string {
	return assemblyVpmullw128
}

func (v *VPMULLW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpmullw128(&vals1, &vals2, &ret)

	log.Printf("VPMULLW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULLW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
