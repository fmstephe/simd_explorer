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
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPAVGB256() *VPAVGB256 {
	return &VPAVGB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPAVGB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPAVGB256) Output() *number.Parameter {
	return v.ret
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

func (v *VPAVGB256) Run(_ [][]byte) (output []byte) {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpavgb256(&vals1, &vals2, &ret)

	log.Printf("VPAVGB256 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *VPAVGB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
