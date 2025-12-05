package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhuw_128.s
var assemblyVpmulhuw128 string

//go:embed stub_vpmulhuw_128.go
var stubVpmulhuw128 string

type VPMULHUW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHUW128() *VPMULHUW128 {
	return &VPMULHUW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPMULHUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHUW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHUW128) Name() string {
	return "VPMULHUW (128 bit) "
}

func (v *VPMULHUW128) Description() string {
	return "Multiply packed unsigned 16-bit integers; keep the high 16 bits of each 32-bit product."
}

func (v *VPMULHUW128) Stub() string {
	return stubVpmulhuw128
}

func (v *VPMULHUW128) Assembly() string {
	return assemblyVpmulhuw128
}

func (v *VPMULHUW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpmulhuw128(&vals1, &vals2, &ret)

	log.Printf("VPMULHUW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMULHUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
