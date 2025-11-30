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
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAVGB128() *VPAVGB128 {
	return &VPAVGB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPAVGB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAVGB128) Output() *number.Parameter {
	return v.ret
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

func (v *VPAVGB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpavgb128(&vals1, &vals2, &ret)

	log.Printf("VPAVGB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)

}

func (v *VPAVGB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
