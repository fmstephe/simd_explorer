package pmaxsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsw_256.s
var assemblyVpmaxsw256 string

//go:embed stub_vpmaxsw_256.go
var stubVpmaxsw256 string

type VPMAXSW256 struct {
}

func (v *VPMAXSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewIntParameter(256, 16, 10),
		number.NewIntParameter(256, 16, 10),
	}
}

func (v *VPMAXSW256) Output() *number.Parameter {
	return number.NewIntParameter(256, 16, 10)
}

func (v *VPMAXSW256) Name() string {
	return "VPMAXSW (256 bit)"
}

func (v *VPMAXSW256) Description() string {
	return "Packed max of signed 16-bit words per lane (VEX, per 128-bit lane)."
}

func (v *VPMAXSW256) Stub() string {
	return stubVpmaxsw256
}

func (v *VPMAXSW256) Assembly() string {
	return assemblyVpmaxsw256
}

func (v *VPMAXSW256) Run(inputs [][]byte) (output []byte) {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(inputs[0]))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(inputs[1]))

	ret := [16]int16{}

	vpmaxsw256(&vals1, &vals2, &ret)

	log.Printf("VPMAXSW256 input %v %v output %v", vals1, vals2, ret)

	return number.Int16SliceToBytes(ret[:])
}

func (v *VPMAXSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
