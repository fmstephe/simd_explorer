package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsw_256.s
var assemblyVpminsw256 string

//go:embed stub_vpminsw_256.go
var stubVpminsw256 string

type VPMINSW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSW256() *VPMINSW256 {
	return &VPMINSW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMINSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSW256) Name() string {
	return "VPMINSW (256 bit) "
}

func (v *VPMINSW256) Description() string {
	return "Signed minimum of packed 16-bit integers."
}

func (v *VPMINSW256) Stub() string {
	return stubVpminsw256
}

func (v *VPMINSW256) Assembly() string {
	return assemblyVpminsw256
}

func (v *VPMINSW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [16]int16{}

	vpminsw256(&vals1, &vals2, &ret)

	log.Printf("VPMINSW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMINSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
