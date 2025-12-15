package vpmaddwd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaddwd_128.s
var assemblyVpmaddwd128 string

//go:embed stub_vpmaddwd_128.go
var stubVpmaddwd128 string

type VPMADDWD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMADDWD128() *VPMADDWD128 {
	return &VPMADDWD128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VPMADDWD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMADDWD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMADDWD128) Name() string {
	return "VPMADDWD (128 bit) "
}

func (v *VPMADDWD128) Description() string {
	return "Multiply adjacent signed 16-bit integers from vals1 and vals2, " +
		"then add each pair of products to produce 32-bit results in ret."
}

func (v *VPMADDWD128) Stub() string {
	return stubVpmaddwd128
}

func (v *VPMADDWD128) Assembly() string {
	return assemblyVpmaddwd128
}

func (v *VPMADDWD128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vpmaddwd128(&vals1, &vals2, &ret)

	log.Printf("VPMADDWD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMADDWD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
