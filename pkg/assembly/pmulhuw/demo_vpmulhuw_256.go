package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhuw_256.s
var assemblyVpmulhuw256 string

//go:embed stub_vpmulhuw_256.go
var stubVpmulhuw256 string

type VPMULHUW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHUW256() *VPMULHUW256 {
	return &VPMULHUW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPMULHUW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHUW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHUW256) Name() string {
	return "VPMULHUW (256 bit)"
}

func (v *VPMULHUW256) Description() string {
	return "Packed multiply unsigned 16-bit integers (VEX, per 128-bit lane); keep the high 16 bits."
}

func (v *VPMULHUW256) Stub() string {
	return stubVpmulhuw256
}

func (v *VPMULHUW256) Assembly() string {
	return assemblyVpmulhuw256
}

func (v *VPMULHUW256) Run() (output []byte) {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpmulhuw256(&vals1, &vals2, &ret)

	log.Printf("VPMULHUW256 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPMULHUW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
