package pmaxs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsd_256.s
var assemblyVpmaxsd256 string

//go:embed stub_vpmaxsd_256.go
var stubVpmaxsd256 string

type VPMAXSD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSD256() *VPMAXSD256 {
	return &VPMAXSD256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMAXSD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSD256) Name() string {
	return "VPMAXSD (256 bit) "
}

func (v *VPMAXSD256) Description() string {
	return "Signed maximum of packed 32-bit integers."
}

func (v *VPMAXSD256) Stub() string {
	return stubVpmaxsd256
}

func (v *VPMAXSD256) Assembly() string {
	return assemblyVpmaxsd256
}

func (v *VPMAXSD256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))

	ret := [8]int32{}

	vpmaxsd256(&vals1, &vals2, &ret)

	log.Printf("VPMAXSD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXSD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
