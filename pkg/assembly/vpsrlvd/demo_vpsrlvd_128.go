package vpsrlvd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlvd_128.s
var assemblyVpsrlvd128 string

//go:embed stub_vpsrlvd_128.go
var stubVpsrlvd128 string

type VPSRLVD128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRLVD128() *VPSRLVD128 {
	return &VPSRLVD128{
		vals:   number.NewNamedUintParameter("vals", 128, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSRLVD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRLVD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLVD128) Name() string {
	return "VPSRLVD (128 bit) "
}

func (v *VPSRLVD128) Description() string {
	return "Shift packed doubleword integers right by variable counts per lane (VEX.128)."
}

func (v *VPSRLVD128) Stub() string {
	return stubVpsrlvd128
}

func (v *VPSRLVD128) Assembly() string {
	return assemblyVpsrlvd128
}

func (v *VPSRLVD128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shifts := [4]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [4]uint32{}

	vpsrlvd128(&vals, &shifts, &ret)

	log.Printf("VPSRLVD128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLVD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
