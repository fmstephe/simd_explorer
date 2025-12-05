package pmulhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhw_128.s
var assemblyVpmulhw128 string

//go:embed stub_vpmulhw_128.go
var stubVpmulhw128 string

type VPMULHW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHW128() *VPMULHW128 {
	return &VPMULHW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMULHW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHW128) Name() string {
	return "VPMULHW (128 bit) "
}

func (v *VPMULHW128) Description() string {
	return "Multiply packed signed 16-bit integers and keep high 16 bits of 32-bit products."
}

func (v *VPMULHW128) Stub() string {
	return stubVpmulhw128
}

func (v *VPMULHW128) Assembly() string {
	return assemblyVpmulhw128
}

func (v *VPMULHW128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	vpmulhw128(&vals1, &vals2, &ret)

	log.Printf("VPMULHW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULHW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
