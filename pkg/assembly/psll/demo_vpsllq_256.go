package psll

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsllq_256.s
var assemblyVpsllq256 string

//go:embed stub_vpsllq_256.go
var stubVpsllq256 string

type VPSLLQ256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSLLQ256() *VPSLLQ256 {
	return &VPSLLQ256{
		vals:  number.NewNamedUintParameter("vals", 256, 64, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPSLLQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSLLQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLQ256) Name() string {
	return "VPSLLQ (256 bit) "
}

func (v *VPSLLQ256) Description() string {
	return "Logical left shift of packed 64-bit integers by per-lane counts from register."
}

func (v *VPSLLQ256) Stub() string {
	return stubVpsllq256
}

func (v *VPSLLQ256) Assembly() string {
	return assemblyVpsllq256
}

func (v *VPSLLQ256) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [4]uint64{}

	vpsllq256(&vals, &shift, &ret)

	log.Printf("VPSLLQ256 vals %v shift %v ret %v", vals, shift, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSLLQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
