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
}

func (v *PMINSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewIntParameter(128, 16, 10),
		number.NewIntParameter(128, 16, 10),
	}
}

func (v *PMINSW128) Output() *number.Parameter {
	return number.NewIntParameter(128, 16, 10)
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

func (v *PMINSW128) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(inputs[0]))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(inputs[1]))

	ret := [8]int16{}

	pminsw128(&vals1, &vals2, &ret)

	log.Printf("PMINSW128 input %v %v output %v", vals1, vals2, ret)

	return number.Int16SliceToBytes(ret[:])
}

func (v *PMINSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
