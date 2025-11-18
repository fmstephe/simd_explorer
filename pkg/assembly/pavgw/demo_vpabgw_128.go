package pavgw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpabgw_128.s
var assemblyVpabgw128 string

//go:embed stub_vpabgw_128.go
var stubVpabgw128 string

type VPAVGW128 struct {
}

func (v *VPAVGW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *VPAVGW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *VPAVGW128) Name() string {
	return "VPAVGW (128 bit)"
}

func (v *VPAVGW128) Description() string {
	return "Average of packed unsigned 16-bit words with rounding (VEX): (a+b+1)>>1."
}

func (v *VPAVGW128) Stub() string {
	return stubVpabgw128
}

func (v *VPAVGW128) Assembly() string {
	return assemblyVpabgw128
}

func (v *VPAVGW128) Run(inputs [][]byte) (output []byte) {
	u1 := [8]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [8]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [8]uint16{}

	vpabgw128(&u1, &u2, &ret)

	log.Printf("VPAVGW128 input %v %v output %v", u1, u2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPAVGW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
