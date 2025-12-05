package pmaddhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddhw_256.s
var assemblyVpmaddhw256 string

//go:embed stub_vpmaddhw_256.go
var stubVpmaddhw256 string

type VPMADDHW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDHW256() *VPMADDHW256 {
	return &VPMADDHW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMADDHW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDHW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDHW256) Name() string {
	return "VPMADDHW (256 bit) "
}

func (v *VPMADDHW256) Description() string {
	return "Multiply signed 16-bit words, add adjacent products, produce 32-bit results."
}

func (v *VPMADDHW256) Stub() string {
	return stubVpmaddhw256
}

func (v *VPMADDHW256) Assembly() string {
	return assemblyVpmaddhw256
}

func (v *VPMADDHW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [8]int32{}

	vpmaddhw256(&vals1, &vals2, &ret)

	log.Printf("VPMADDHW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMADDHW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
