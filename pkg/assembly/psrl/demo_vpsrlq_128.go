package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlq_128.s
var assemblyVpsrlq128 string

//go:embed stub_vpsrlq_128.go
var stubVpsrlq128 string

type VPSRLQ128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLQ128() *VPSRLQ128 {
	return &VPSRLQ128{
		vals:  number.NewNamedUintParameter("vals", 128, 64, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPSRLQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLQ128) Name() string {
	return "VPSRLQ (128 bit) "
}

func (v *VPSRLQ128) Description() string {
	return "Logical right shift of packed 64-bit integers by per-lane counts from register."
}

func (v *VPSRLQ128) Stub() string {
	return stubVpsrlq128
}

func (v *VPSRLQ128) Assembly() string {
	return assemblyVpsrlq128
}

func (v *VPSRLQ128) Run() {
	vals := [2]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [2]uint64{}

	vpsrlq128(&vals, &shift, &ret)

	log.Printf("VPSRLQ128 vals %v shift %v ret %v", vals, shift, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
