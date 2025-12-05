package pmulhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhw_256.s
var assemblyVpmulhw256 string

//go:embed stub_vpmulhw_256.go
var stubVpmulhw256 string

type VPMULHW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHW256() *VPMULHW256 {
	return &VPMULHW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMULHW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHW256) Name() string {
	return "VPMULHW (256 bit) "
}

func (v *VPMULHW256) Description() string {
	return "Multiply packed signed 16-bit integers and keep high 16 bits of 32-bit products."
}

func (v *VPMULHW256) Stub() string {
	return stubVpmulhw256
}

func (v *VPMULHW256) Assembly() string {
	return assemblyVpmulhw256
}

func (v *VPMULHW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [16]int16{}

	vpmulhw256(&vals1, &vals2, &ret)

	log.Printf("VPMULHW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULHW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
