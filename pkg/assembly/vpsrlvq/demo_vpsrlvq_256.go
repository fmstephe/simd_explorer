package vpsrlvq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlvq_256.s
var assemblyVpsrlvq256 string

//go:embed stub_vpsrlvq_256.go
var stubVpsrlvq256 string

type VPSRLVQ256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRLVQ256() *VPSRLVQ256 {
	return &VPSRLVQ256{
		vals:   number.NewNamedUintParameter("vals", 256, 64, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 64, 10),
		ret:    number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPSRLVQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRLVQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLVQ256) Name() string {
	return "VPSRLVQ (256 bit) "
}

func (v *VPSRLVQ256) Description() string {
	return "Shift packed quadword integers right by variable counts per lane (VEX.256)."
}

func (v *VPSRLVQ256) Stub() string {
	return stubVpsrlvq256
}

func (v *VPSRLVQ256) Assembly() string {
	return assemblyVpsrlvq256
}

func (v *VPSRLVQ256) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shifts := [4]uint64{}
	copy(shifts[:], number.ToUint64Slice(v.shifts.FlatData()))

	ret := [4]uint64{}

	vpsrlvq256(&vals, &shifts, &ret)

	log.Printf("VPSRLVQ256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLVQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
