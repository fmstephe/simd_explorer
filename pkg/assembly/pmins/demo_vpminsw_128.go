package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsw_128.s
var assemblyVpminsw128 string

//go:embed stub_vpminsw_128.go
var stubVpminsw128 string

type VPMINSW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSW128() *VPMINSW128 {
	return &VPMINSW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMINSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSW128) Name() string {
	return "VPMINSW (128 bit) "
}

func (v *VPMINSW128) Description() string {
	return "Signed minimum of packed 16-bit integers."
}

func (v *VPMINSW128) Stub() string {
	return stubVpminsw128
}

func (v *VPMINSW128) Assembly() string {
	return assemblyVpminsw128
}

func (v *VPMINSW128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	vpminsw128(&vals1, &vals2, &ret)

	log.Printf("VPMINSW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
