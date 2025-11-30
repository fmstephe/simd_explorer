package pminub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminub_128.s
var assemblyVpminub128 string

//go:embed stub_vpminub_128.go
var stubVpminub128 string

type VPMINUB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINUB128() *VPMINUB128 {
	return &VPMINUB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPMINUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINUB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINUB128) Name() string {
	return "VPMINUB (128 bit)"
}

func (v *VPMINUB128) Description() string {
	return "Packed min of unsigned bytes per lane (VEX)."
}

func (v *VPMINUB128) Stub() string {
	return stubVpminub128
}

func (v *VPMINUB128) Assembly() string {
	return assemblyVpminub128
}

func (v *VPMINUB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpminub128(&vals1, &vals2, &ret)

	log.Printf("VPMINUB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)

}

func (v *VPMINUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
