package pminub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminub_256.s
var assemblyVpminub256 string

//go:embed stub_vpminub_256.go
var stubVpminub256 string

type VPMINUB256 struct {
}

func (v *VPMINUB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 8, 10),
		number.NewUintParameter(256, 8, 10),
	}
}

func (v *VPMINUB256) Output() *number.Parameter {
	return number.NewUintParameter(256, 8, 10)
}

func (v *VPMINUB256) Name() string {
	return "VPMINUB (256 bit)"
}

func (v *VPMINUB256) Description() string {
	return "Packed min of unsigned bytes per lane (VEX, per 128-bit lane)."
}

func (v *VPMINUB256) Stub() string {
	return stubVpminub256
}

func (v *VPMINUB256) Assembly() string {
	return assemblyVpminub256
}

func (v *VPMINUB256) Run(inputs [][]byte) (output []byte) {
	b1 := [32]uint8{}
	copy(b1[:], inputs[0])
	b2 := [32]uint8{}
	copy(b2[:], inputs[1])

	ret := [32]uint8{}

	vpminub256(&b1, &b2, &ret)

	log.Printf("VPMINUB256 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPMINUB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
