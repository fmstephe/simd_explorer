package vpsllvq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvq_128.s
var assemblyVpsllvq128 string

//go:embed stub_vpsllvq_128.go
var stubVpsllvq128 string

type VPSLLVQ128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVQ128() *VPSLLVQ128 {
	return &VPSLLVQ128{
		vals:   number.NewNamedUintParameter("vals", 128, 64, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 64, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPSLLVQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVQ128) Name() string {
	return "VPSLLVQ (128 bit) "
}

func (v *VPSLLVQ128) Description() string {
	return "Shift packed quadword integers left by variable counts per lane (VEX.128)."
}

func (v *VPSLLVQ128) Stub() string {
	return stubVpsllvq128
}

func (v *VPSLLVQ128) Assembly() string {
	return assemblyVpsllvq128
}

func (v *VPSLLVQ128) Run() {
	vals := [2]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shifts := [2]uint64{}
	copy(shifts[:], number.ToUint64Slice(v.shifts.FlatData()))

	ret := [2]uint64{}

	vpsllvq128(&vals, &shifts, &ret)

	log.Printf("VPSLLVQ128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
