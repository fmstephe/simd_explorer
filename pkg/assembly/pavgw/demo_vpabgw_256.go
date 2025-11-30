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
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAVGW256() *VPAVGW256 {
	return &VPAVGW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPAVGW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAVGW256) Output() *number.Parameter {
	return v.ret
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

func (v *VPAVGW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpabgw256(&vals1, &vals2, &ret)

	log.Printf("VPAVGW256 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPAVGW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
