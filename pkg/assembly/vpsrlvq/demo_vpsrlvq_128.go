package vpsrlvq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlvq_128.s
var assemblyVpsrlvq128 string

//go:embed stub_vpsrlvq_128.go
var stubVpsrlvq128 string

type VPSRLVQ128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRLVQ128() *VPSRLVQ128 {
	return &VPSRLVQ128{
		vals:   number.NewNamedUintParameter("vals", 128, 64, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 64, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPSRLVQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRLVQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLVQ128) Name() string {
	return "VPSRLVQ (128 bit) "
}

func (v *VPSRLVQ128) Description() string {
	return "Shift packed quadword integers right by variable counts per lane (VEX.128)."
}

func (v *VPSRLVQ128) Stub() string {
	return stubVpsrlvq128
}

func (v *VPSRLVQ128) Assembly() string {
	return assemblyVpsrlvq128
}

func (v *VPSRLVQ128) Run() {
	vals := [2]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shifts := [2]uint64{}
	copy(shifts[:], number.ToUint64Slice(v.shifts.FlatData()))

	ret := [2]uint64{}

	vpsrlvq128(&vals, &shifts, &ret)

	log.Printf("VPSRLVQ128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLVQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
