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
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUB256() *VPMINUB256 {
	return &VPMINUB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPMINUB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUB256) Output() *number.Parameter {
	return v.ret
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

func (v *VPMINUB256) Run(_ [][]byte) (output []byte) {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpminub256(&vals1, &vals2, &ret)

	log.Printf("VPMINUB256 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *VPMINUB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
