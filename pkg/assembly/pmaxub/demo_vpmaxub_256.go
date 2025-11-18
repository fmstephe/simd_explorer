package pmaxub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxub_256.s
var assemblyVpmaxub256 string

//go:embed stub_vpmaxub_256.go
var stubVpmaxub256 string

type VPMAXUB256 struct {
}

func (v *VPMAXUB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 8, 10),
		number.NewUintParameter(256, 8, 10),
	}
}

func (v *VPMAXUB256) Output() *number.Parameter {
	return number.NewUintParameter(256, 8, 10)
}

func (v *VPMAXUB256) Name() string {
	return "VPMAXUB (256 bit)"
}

func (v *VPMAXUB256) Description() string {
	return "Packed max of unsigned bytes per lane (VEX, per 128-bit lane)."
}

func (v *VPMAXUB256) Stub() string {
	return stubVpmaxub256
}

func (v *VPMAXUB256) Assembly() string {
	return assemblyVpmaxub256
}

func (v *VPMAXUB256) Run(inputs [][]byte) (output []byte) {
	b1 := [32]uint8{}
	copy(b1[:], inputs[0])
	b2 := [32]uint8{}
	copy(b2[:], inputs[1])

	ret := [32]uint8{}

	vpmaxub256(&b1, &b2, &ret)

	log.Printf("VPMAXUB256 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPMAXUB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
