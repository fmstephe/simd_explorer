package pavgw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pabgw_128.s
var assemblyPabgw128 string

//go:embed stub_pabgw_128.go
var stubPabgw128 string

type PAVGW128 struct {
}

func (v *PAVGW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *PAVGW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PAVGW128) Name() string {
	return "PAVGW (128 bit)"
}

func (v *PAVGW128) Description() string {
	return "Average of packed unsigned 16-bit words with rounding: (a+b+1)>>1."
}

func (v *PAVGW128) Stub() string {
	return stubPabgw128
}

func (v *PAVGW128) Assembly() string {
	return assemblyPabgw128
}

func (v *PAVGW128) Run(inputs [][]byte) (output []byte) {
	u1 := [8]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [8]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [8]uint16{}

	pabgw128(&u1, &u2, &ret)

	log.Printf("PAVGW128 input %v %v output %v", u1, u2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PAVGW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
