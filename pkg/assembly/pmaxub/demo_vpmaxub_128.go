package pmaxub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxub_128.s
var assemblyVpmaxub128 string

//go:embed stub_vpmaxub_128.go
var stubVpmaxub128 string

type VPMAXUB128 struct {
}

func (v *VPMAXUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *VPMAXUB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *VPMAXUB128) Name() string {
	return "VPMAXUB (128 bit)"
}

func (v *VPMAXUB128) Description() string {
	return "Packed max of unsigned bytes per lane (VEX)."
}

func (v *VPMAXUB128) Stub() string {
	return stubVpmaxub128
}

func (v *VPMAXUB128) Assembly() string {
	return assemblyVpmaxub128
}

func (v *VPMAXUB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	vpmaxub128(&b1, &b2, &ret)

	log.Printf("VPMAXUB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPMAXUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
