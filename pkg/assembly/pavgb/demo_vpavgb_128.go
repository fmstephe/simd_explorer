package pavgb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpavgb_128.s
var assemblyVpavgb128 string

//go:embed stub_vpavgb_128.go
var stubVpavgb128 string

type VPAVGB128 struct {
}

func (v *VPAVGB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *VPAVGB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *VPAVGB128) Name() string {
	return "VPAVGB (128 bit)"
}

func (v *VPAVGB128) Description() string {
	return "Average of packed unsigned bytes with rounding (VEX): (a+b+1)>>1."
}

func (v *VPAVGB128) Stub() string {
	return stubVpavgb128
}

func (v *VPAVGB128) Assembly() string {
	return assemblyVpavgb128
}

func (v *VPAVGB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	vpavgb128(&b1, &b2, &ret)

	log.Printf("VPAVGB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *VPAVGB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
