package pmaxub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmaxub_128.s
var assemblyPmaxub128 string

//go:embed stub_pmaxub_128.go
var stubPmaxub128 string

type PMAXUB128 struct {
}

func (v *PMAXUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *PMAXUB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *PMAXUB128) Name() string {
	return "PMAXUB (128 bit)"
}

func (v *PMAXUB128) Description() string {
	return "Packed max of unsigned bytes per lane."
}

func (v *PMAXUB128) Stub() string {
	return stubPmaxub128
}

func (v *PMAXUB128) Assembly() string {
	return assemblyPmaxub128
}

func (v *PMAXUB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	pmaxub128(&b1, &b2, &ret)

	log.Printf("PMAXUB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *PMAXUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
