package pminub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pminub_128.s
var assemblyPminub128 string

//go:embed stub_pminub_128.go
var stubPminub128 string

type PMINUB128 struct {
}

func (v *PMINUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *PMINUB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *PMINUB128) Name() string {
	return "PMINUB (128 bit)"
}

func (v *PMINUB128) Description() string {
	return "Packed min of unsigned bytes per lane."
}

func (v *PMINUB128) Stub() string {
	return stubPminub128
}

func (v *PMINUB128) Assembly() string {
	return assemblyPminub128
}

func (v *PMINUB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	pminub128(&b1, &b2, &ret)

	log.Printf("PMINUB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *PMINUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
