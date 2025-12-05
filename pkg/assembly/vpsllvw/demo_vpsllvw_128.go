package vpsllvw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvw_128.s
var assemblyVpsllvw128 string

//go:embed stub_vpsllvw_128.go
var stubVpsllvw128 string

type VPSLLVW128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVW128() *VPSLLVW128 {
	return &VPSLLVW128{
		vals:   number.NewNamedUintParameter("vals", 128, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSLLVW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVW128) Name() string {
	return "VPSLLVW (128 bit) "
}

func (v *VPSLLVW128) Description() string {
	return "Shift packed doubleword integers left by variable counts per lane (VEX.128)."
}

func (v *VPSLLVW128) Stub() string {
	return stubVpsllvw128
}

func (v *VPSLLVW128) Assembly() string {
	return assemblyVpsllvw128
}

func (v *VPSLLVW128) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	shifts := [8]uint16{}
	copy(shifts[:], number.ToUint16Slice(v.shifts.FlatData()))

	ret := [8]uint16{}

	vpsllvw128(&vals, &shifts, &ret)

	log.Printf("VPSLLVW128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
