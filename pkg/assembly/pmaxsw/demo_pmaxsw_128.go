package pmaxsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmaxsw_128.s
var assemblyPmaxsw128 string

//go:embed stub_pmaxsw_128.go
var stubPmaxsw128 string

type PMAXSW128 struct {
}

func (v *PMAXSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewIntParameter(128, 16, 10),
		number.NewIntParameter(128, 16, 10),
	}
}

func (v *PMAXSW128) Output() *number.Parameter {
	return number.NewIntParameter(128, 16, 10)
}

func (v *PMAXSW128) Name() string {
	return "PMAXSW (128 bit)"
}

func (v *PMAXSW128) Description() string {
	return "Packed max of signed 16-bit words per lane."
}

func (v *PMAXSW128) Stub() string {
	return stubPmaxsw128
}

func (v *PMAXSW128) Assembly() string {
	return assemblyPmaxsw128
}

func (v *PMAXSW128) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(inputs[0]))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(inputs[1]))

	ret := [8]int16{}

	pmaxsw128(&vals1, &vals2, &ret)

	log.Printf("PMAXSW128 input %v %v output %v", vals1, vals2, ret)

	return number.Int16SliceToBytes(ret[:])
}

func (v *PMAXSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
