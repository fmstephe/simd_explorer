package psra

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsraq_128.s
var assemblyVpsraq128 string

//go:embed stub_vpsraq_128.go
var stubVpsraq128 string

type VPSRAQ128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRAQ128() *VPSRAQ128 {
	return &VPSRAQ128{
		vals:  number.NewNamedIntParameter("vals", 128, 64, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 64, 10),
	}
}

func (v *VPSRAQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRAQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRAQ128) Name() string {
	return "VPSRAQ (128 bit) "
}

func (v *VPSRAQ128) Description() string {
	return "Arithmetical right shift of packed 64-bit integers by per-lane counts from register."
}

func (v *VPSRAQ128) Stub() string {
	return stubVpsraq128
}

func (v *VPSRAQ128) Assembly() string {
	return assemblyVpsraq128
}

func (v *VPSRAQ128) Run() {
	vals := [2]int64{}
	copy(vals[:], number.ToInt64Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [2]int64{}

	vpsraq128(&vals, &shift, &ret)

	log.Printf("VPSRAQ128 vals %v shift %v ret %v", vals, shift, ret)

	out := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRAQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
