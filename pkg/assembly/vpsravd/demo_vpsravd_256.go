package vpsravd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsravd_256.s
var assemblyVpsravd256 string

//go:embed stub_vpsravd_256.go
var stubVpsravd256 string

type VPSRAVD256 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRAVD256() *VPSRAVD256 {
	return &VPSRAVD256{
		vals:   number.NewNamedIntParameter("vals", 256, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 256, 32, 10),
		ret:    number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPSRAVD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRAVD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAVD256) Name() string {
	return "VPSRAVD (256 bit) "
}

func (v *VPSRAVD256) Description() string {
	return "Arithmetic right shift packed doubleword integers by variable counts per lane (VEX.256)."
}

func (v *VPSRAVD256) Stub() string {
	return stubVpsravd256
}

func (v *VPSRAVD256) Assembly() string {
	return assemblyVpsravd256
}

func (v *VPSRAVD256) Run() {
	vals := [8]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	shifts := [8]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [8]int32{}

	vpsravd256(&vals, &shifts, &ret)

	log.Printf("VPSRAVD256 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAVD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
