package pshufb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufb_128.s
var assemblyVpshufb128 string

//go:embed stub_vpshufb_128.go
var stubVpshufb128 string

type VPSHUFB128 struct {
	vals1   *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPSHUFB128() *VPSHUFB128 {
	return &VPSHUFB128{
		vals1:   number.NewNamedUintParameter("vals1", 128, 8, 10),
		control: number.NewNamedUintParameter("control", 128, 8, 10),
		ret:     number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSHUFB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.control,
	}
}

func (v *VPSHUFB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFB128) Name() string {
	return "VPSHUFB (128 bit) "
}

func (v *VPSHUFB128) Description() string {
	return "Shuffle bytes in vals1 according to control; high bit in control zeroes out."
}

func (v *VPSHUFB128) Stub() string {
	return stubVpshufb128
}

func (v *VPSHUFB128) Assembly() string {
	return assemblyVpshufb128
}

func (v *VPSHUFB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	control := [16]uint8{}
	copy(control[:], v.control.FlatData())

	ret := [16]uint8{}

	vpshufb128(&vals1, &control, &ret)

	log.Printf("VPSHUFB128 vals1 %v control %v ret %v", vals1, control, ret)

	v.ret.SetData(ret[:])
}

func (v *VPSHUFB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
