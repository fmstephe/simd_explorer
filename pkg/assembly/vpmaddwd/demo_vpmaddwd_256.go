package vpmaddwd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddwd_256.s
var assemblyVpmaddwd256 string

//go:embed stub_vpmaddwd_256.go
var stubVpmaddwd256 string

type VPMADDWD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDWD256() *VPMADDWD256 {
	return &VPMADDWD256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VPMADDWD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDWD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDWD256) Name() string {
	return "VPMADDWD (256 bit) "
}

func (v *VPMADDWD256) Description() string {
	return "Multiply adjacent signed 16-bit integers from vals1 and vals2, " +
		"then add each pair of products to produce 32-bit results in ret (per 128-bit lane)."
}

func (v *VPMADDWD256) Stub() string {
	return stubVpmaddwd256
}

func (v *VPMADDWD256) Assembly() string {
	return assemblyVpmaddwd256
}

func (v *VPMADDWD256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmaddwd256(&vals1, &vals2, &ret)

	log.Printf("VPMADDWD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMADDWD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
