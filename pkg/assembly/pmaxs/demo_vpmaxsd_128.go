package pmaxs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsd_128.s
var assemblyVpmaxsd128 string

//go:embed stub_vpmaxsd_128.go
var stubVpmaxsd128 string

type VPMAXSD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSD128() *VPMAXSD128 {
	return &VPMAXSD128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMAXSD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSD128) Name() string {
	return "VPMAXSD (128 bit) "
}

func (v *VPMAXSD128) Description() string {
	return "Signed maximum of packed 32-bit integers."
}

func (v *VPMAXSD128) Stub() string {
	return stubVpmaxsd128
}

func (v *VPMAXSD128) Assembly() string {
	return assemblyVpmaxsd128
}

func (v *VPMAXSD128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [4]int32{}

	vpmaxsd128(&vals1, &vals2, &ret)

	log.Printf("VPMAXSD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXSD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
