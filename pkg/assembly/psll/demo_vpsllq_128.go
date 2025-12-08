package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllq_128.s
var assemblyVpsllq128 string

//go:embed stub_vpsllq_128.go
var stubVpsllq128 string

type VPSLLQ128 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLQ128() *VPSLLQ128 {
	return &VPSLLQ128{
		vals:  number.NewNamedUintParameter("vals", 128, 64, 10),
		shift: number.NewNamedUintParameter("shift", 128, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPSLLQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLQ128) Name() string {
	return "VPSLLQ (128 bit) "
}

func (v *VPSLLQ128) Description() string {
	return "Logical left shift of packed 64-bit integers by per-lane counts from register."
}

func (v *VPSLLQ128) Stub() string {
	return stubVpsllq128
}

func (v *VPSLLQ128) Assembly() string {
	return assemblyVpsllq128
}

func (v *VPSLLQ128) Run() {
	vals := [2]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shift := [2]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [2]uint64{}

	vpsllq128(&vals, &shift, &ret)

	log.Printf("VPSLLQ128 vals %v shift %v ret %v", vals, shift, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
