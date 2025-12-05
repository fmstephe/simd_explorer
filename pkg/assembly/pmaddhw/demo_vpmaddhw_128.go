package pmaddhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddhw_128.s
var assemblyVpmaddhw128 string

//go:embed stub_vpmaddhw_128.go
var stubVpmaddhw128 string

type VPMADDHW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDHW128() *VPMADDHW128 {
	return &VPMADDHW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMADDHW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDHW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDHW128) Name() string {
	return "VPMADDHW (128 bit) "
}

func (v *VPMADDHW128) Description() string {
	return "Multiply signed 16-bit words, add adjacent products, produce 32-bit results."
}

func (v *VPMADDHW128) Stub() string {
	return stubVpmaddhw128
}

func (v *VPMADDHW128) Assembly() string {
	return assemblyVpmaddhw128
}

func (v *VPMADDHW128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [4]int32{}

	vpmaddhw128(&vals1, &vals2, &ret)

	log.Printf("VPMADDHW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMADDHW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
