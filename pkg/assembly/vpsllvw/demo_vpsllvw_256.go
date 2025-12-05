package vpsllvw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvw_256.s
var assemblyVpsllvw256 string

//go:embed stub_vpsllvw_256.go
var stubVpsllvw256 string

type VPSLLVW256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVW256() *VPSLLVW256 {
	return &VPSLLVW256{
		vals:   number.NewNamedUintParameter("vals", 256, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSLLVW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVW256) Name() string {
	return "VPSLLVW (256 bit) "
}

func (v *VPSLLVW256) Description() string {
	return "Shift packed doubleword integers left by variable counts per lane (VEX.256)."
}

func (v *VPSLLVW256) Stub() string {
	return stubVpsllvw256
}

func (v *VPSLLVW256) Assembly() string {
	return assemblyVpsllvw256
}

func (v *VPSLLVW256) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shifts := [16]uint16{}
	copy(shifts[:], number.ToUint16Slice(v.shifts.FlatData()))

	ret := [16]uint16{}

	vpsllvw256(&vals, &shifts, &ret)

	log.Printf("VPSLLVW256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
