package pminsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsw_128.s
var assemblyVpminsw128 string

//go:embed stub_vpminsw_128.go
var stubVpminsw128 string

type VPMINSW128 struct {
}

func (v *VPMINSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewIntParameter(128, 16, 10),
		number.NewIntParameter(128, 16, 10),
	}
}

func (v *VPMINSW128) Output() *number.Parameter {
	return number.NewIntParameter(128, 16, 10)
}

func (v *VPMINSW128) Name() string {
	return "VPMINSW (128 bit)"
}

func (v *VPMINSW128) Description() string {
	return "Packed min of signed 16-bit words per lane (VEX)."
}

func (v *VPMINSW128) Stub() string {
	return stubVpminsw128
}

func (v *VPMINSW128) Assembly() string {
	return assemblyVpminsw128
}

func (v *VPMINSW128) Run(inputs [][]byte) (output []byte) {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(inputs[0]))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(inputs[1]))

	ret := [8]int16{}

	vpminsw128(&vals1, &vals2, &ret)

	log.Printf("VPMINSW128 input %v %v output %v", vals1, vals2, ret)

	return number.Int16SliceToBytes(ret[:])
}

func (v *VPMINSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
