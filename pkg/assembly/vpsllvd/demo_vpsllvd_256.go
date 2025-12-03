package vpsllvd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvd_256.s
var assemblyVpsllvd256 string

//go:embed stub_vpsllvd_256.go
var stubVpsllvd256 string

type VPSLLVD256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVD256() *VPSLLVD256 {
	return &VPSLLVD256{
		vals:   number.NewNamedUintParameter("vals", 256, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSLLVD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVD256) Name() string {
	return "VPSLLVD (256 bit) "
}

func (v *VPSLLVD256) Description() string {
	return "Shift packed doubleword integers left by variable counts per lane (VEX.256)."
}

func (v *VPSLLVD256) Stub() string {
	return stubVpsllvd256
}

func (v *VPSLLVD256) Assembly() string {
	return assemblyVpsllvd256
}

func (v *VPSLLVD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shifts := [8]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [8]uint32{}

	vpsllvd256(&vals, &shifts, &ret)

	log.Printf("VPSLLVD256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
