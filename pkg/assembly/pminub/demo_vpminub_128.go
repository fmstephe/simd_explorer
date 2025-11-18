package pminub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminub_128.s
var assemblyVpminub128 string

//go:embed stub_vpminub_128.go
var stubVpminub128 string

type VPMINUB128 struct {
}

func (v *VPMINUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *VPMINUB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *VPMINUB128) Name() string {
	return "VPMINUB (128 bit)"
}

func (v *VPMINUB128) Description() string {
	return "Packed min of unsigned bytes per lane (VEX)."
}

func (v *VPMINUB128) Stub() string {
	return stubVpminub128
}

func (v *VPMINUB128) Assembly() string {
	return assemblyVpminub128
}

func (v *VPMINUB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	vpminub128(&b1, &b2, &ret)

	log.Printf("VPMINUB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPMINUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
