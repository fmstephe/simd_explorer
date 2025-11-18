package pavgw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabgw_256.s
var assemblyVpabgw256 string

//go:embed stub_vpabgw_256.go
var stubVpabgw256 string

type VPAVGW256 struct {
}

func (v *VPAVGW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 16, 10),
		number.NewUintParameter(256, 16, 10),
	}
}

func (v *VPAVGW256) Output() *number.Parameter {
	return number.NewUintParameter(256, 16, 10)
}

func (v *VPAVGW256) Name() string {
	return "VPAVGW (256 bit)"
}

func (v *VPAVGW256) Description() string {
	return "Average of packed unsigned 16-bit words with rounding (VEX, per 128-bit lane)."
}

func (v *VPAVGW256) Stub() string {
	return stubVpabgw256
}

func (v *VPAVGW256) Assembly() string {
	return assemblyVpabgw256
}

func (v *VPAVGW256) Run(inputs [][]byte) (output []byte) {
	u1 := [16]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [16]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [16]uint16{}

	vpabgw256(&u1, &u2, &ret)

	log.Printf("VPAVGW256 input %v %v output %v", u1, u2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPAVGW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
