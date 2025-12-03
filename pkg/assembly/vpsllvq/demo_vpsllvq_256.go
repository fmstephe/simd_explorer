package vpsllvq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllvq_256.s
var assemblyVpsllvq256 string

//go:embed stub_vpsllvq_256.go
var stubVpsllvq256 string

type VPSLLVQ256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSLLVQ256() *VPSLLVQ256 {
	return &VPSLLVQ256{
		vals:   number.NewNamedUintParameter("vals", 256, 64, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 64, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPSLLVQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSLLVQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLVQ256) Name() string {
	return "VPSLLVQ (256 bit) "
}

func (v *VPSLLVQ256) Description() string {
	return "Shift packed quadword integers left by variable counts per lane (VEX.256)."
}

func (v *VPSLLVQ256) Stub() string {
	return stubVpsllvq256
}

func (v *VPSLLVQ256) Assembly() string {
	return assemblyVpsllvq256
}

func (v *VPSLLVQ256) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shifts := [4]uint64{}
	copy(shifts[:], number.ToUint64Slice(v.shifts.FlatData()))

	ret := [4]uint64{}

	vpsllvq256(&vals, &shifts, &ret)

	log.Printf("VPSLLVQ256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLVQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
