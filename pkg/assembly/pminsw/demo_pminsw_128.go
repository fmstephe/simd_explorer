package pminsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pminsw_128.s
var assemblyPminsw128 string

//go:embed stub_pminsw_128.go
var stubPminsw128 string

type PMINSW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPMINSW128() *PMINSW128 {
	return &PMINSW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *PMINSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PMINSW128) Output() *number.Parameter {
	return v.ret
}

func (v *PMINSW128) Name() string {
	return "PMINSW (128 bit)"
}

func (v *PMINSW128) Description() string {
	return "Packed min of signed 16-bit words per lane."
}

func (v *PMINSW128) Stub() string {
	return stubPminsw128
}

func (v *PMINSW128) Assembly() string {
	return assemblyPminsw128
}

func (v *PMINSW128) Run(_ [][]byte) (output []byte) {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	pminsw128(&vals1, &vals2, &ret)

	log.Printf("PMINSW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *PMINSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
