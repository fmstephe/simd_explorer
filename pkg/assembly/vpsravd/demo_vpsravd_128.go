package vpsravd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsravd_128.s
var assemblyVpsravd128 string

//go:embed stub_vpsravd_128.go
var stubVpsravd128 string

type VPSRAVD128 struct {
	vals   *number.Parameter
	shifts *number.Parameter
	ret    *number.Parameter
}

func NewVPSRAVD128() *VPSRAVD128 {
	return &VPSRAVD128{
		vals:   number.NewNamedIntParameter("vals", 128, 32, 10),
		shifts: number.NewNamedUintParameter("shifts", 128, 32, 10),
		ret:    number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPSRAVD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shifts,
	}
}

func (v *VPSRAVD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAVD128) Name() string {
	return "VPSRAVD (128 bit) "
}

func (v *VPSRAVD128) Description() string {
	return "Arithmetic right shift packed doubleword integers by variable counts per lane (VEX.128)."
}

func (v *VPSRAVD128) Stub() string {
	return stubVpsravd128
}

func (v *VPSRAVD128) Assembly() string {
	return assemblyVpsravd128
}

func (v *VPSRAVD128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	shifts := [4]uint32{}
	copy(shifts[:], number.ToUint32Slice(v.shifts.FlatData()))

	ret := [4]int32{}

	vpsravd128(&vals, &shifts, &ret)

	log.Printf("VPSRAVD128 vals %v shifts %v ret %v", vals, shifts, ret)

	out := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAVD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
