package vpsrlvd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlvd_256.s
var assemblyVpsrlvd256 string

//go:embed stub_vpsrlvd_256.go
var stubVpsrlvd256 string

type VPSRLVD256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRLVD256() *VPSRLVD256 {
	return &VPSRLVD256{
		vals:   number.NewNamedUintParameter("vals", 256, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSRLVD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRLVD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLVD256) Name() string {
	return "VPSRLVD (256 bit) "
}

func (v *VPSRLVD256) Description() string {
	return "Shift packed doubleword integers right by variable counts per lane (VEX.256)."
}

func (v *VPSRLVD256) Stub() string {
	return stubVpsrlvd256
}

func (v *VPSRLVD256) Assembly() string {
	return assemblyVpsrlvd256
}

func (v *VPSRLVD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shifts := [8]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [8]uint32{}

	vpsrlvd256(&vals, &shifts, &ret)

	log.Printf("VPSRLVD256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLVD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
