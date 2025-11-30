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
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAVGW128() *VPAVGW128 {
	return &VPAVGW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPAVGW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAVGW128) Output() *number.Parameter {
	return v.ret
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

func (v *VPAVGW128) Run() (output []byte) {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpabgw128(&vals1, &vals2, &ret)

	log.Printf("VPAVGW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPAVGW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
