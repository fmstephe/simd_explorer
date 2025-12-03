package vpsllvd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvd_128.s
var assemblyVpsllvd128 string

//go:embed stub_vpsllvd_128.go
var stubVpsllvd128 string

type VPSLLVD128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVD128() *VPSLLVD128 {
	return &VPSLLVD128{
		vals:   number.NewNamedUintParameter("vals", 128, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 32, 10),
		ret:    number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSLLVD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVD128) Name() string {
	return "VPSLLVD (128 bit) "
}

func (v *VPSLLVD128) Description() string {
	return "Shift packed doubleword integers left by variable counts per lane (VEX.128)."
}

func (v *VPSLLVD128) Stub() string {
	return stubVpsllvd128
}

func (v *VPSLLVD128) Assembly() string {
	return assemblyVpsllvd128
}

func (v *VPSLLVD128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	shifts := [4]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [4]uint32{}

	vpsllvd128(&vals, &shifts, &ret)

	log.Printf("VPSLLVD128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
