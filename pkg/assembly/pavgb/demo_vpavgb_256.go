package pavgb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpavgb_256.s
var assemblyVpavgb256 string

//go:embed stub_vpavgb_256.go
var stubVpavgb256 string

type VPAVGB256 struct {
}

func (v *VPAVGB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 8, 10),
		number.NewUintParameter(256, 8, 10),
	}
}

func (v *VPAVGB256) Output() *number.Parameter {
	return number.NewUintParameter(256, 8, 10)
}

func (v *VPAVGB256) Name() string {
	return "VPAVGB (256 bit)"
}

func (v *VPAVGB256) Description() string {
	return "Average of packed unsigned bytes with rounding (VEX, per 128-bit lane): (a+b+1)>>1."
}

func (v *VPAVGB256) Stub() string {
	return stubVpavgb256
}

func (v *VPAVGB256) Assembly() string {
	return assemblyVpavgb256
}

func (v *VPAVGB256) Run(inputs [][]byte) (output []byte) {
	b1 := [32]uint8{}
	copy(b1[:], inputs[0])
	b2 := [32]uint8{}
	copy(b2[:], inputs[1])

	ret := [32]uint8{}

	vpavgb256(&b1, &b2, &ret)

	log.Printf("VPAVGB256 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPAVGB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
