package psrl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrlq_256.s
var assemblyVpsrlq256 string

//go:embed stub_vpsrlq_256.go
var stubVpsrlq256 string

type VPSRLQ256 struct {
	vals  *number.Parameter
	shift *number.Parameter
	ret   *number.Parameter
}

func NewVPSRLQ256() *VPSRLQ256 {
	return &VPSRLQ256{
		vals:  number.NewNamedUintParameter("vals", 256, 64, 10),
		shift: number.NewNamedUintParameter("shift", 256, 64, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPSRLQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.shift,
	}
}

func (v *VPSRLQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLQ256) Name() string {
	return "VPSRLQ (256 bit) "
}

func (v *VPSRLQ256) Description() string {
	return "Logical right shift of packed 64-bit integers by per-lane counts from register."
}

func (v *VPSRLQ256) Stub() string {
	return stubVpsrlq256
}

func (v *VPSRLQ256) Assembly() string {
	return assemblyVpsrlq256
}

func (v *VPSRLQ256) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	shift := [4]uint64{}
	copy(shift[:], number.ToUint64Slice(v.shift.FlatData()))

	ret := [4]uint64{}

	vpsrlq256(&vals, &shift, &ret)

	log.Printf("VPSRLQ256 vals %v shift %v ret %v", vals, shift, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSRLQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
