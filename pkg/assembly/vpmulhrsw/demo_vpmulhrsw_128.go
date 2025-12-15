package vpmulhrsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhrsw_128.s
var assemblyVpmulhrsw128 string

//go:embed stub_vpmulhrsw_128.go
var stubVpmulhrsw128 string

type VPMULHRSW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHRSW128() *VPMULHRSW128 {
	return &VPMULHRSW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMULHRSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHRSW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHRSW128) Name() string {
	return "VPMULHRSW (128 bit) "
}

func (v *VPMULHRSW128) Description() string {
	return "Multiply high with round and scale (signed 16→16). " +
		"For each lane, compute (a*b + 0x8000) >> 16 (signed) and store 16-bit results in ret."
}

func (v *VPMULHRSW128) Stub() string {
	return stubVpmulhrsw128
}

func (v *VPMULHRSW128) Assembly() string {
	return assemblyVpmulhrsw128
}

func (v *VPMULHRSW128) Run() {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [8]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpmulhrsw128(&vals1, &vals2, &ret)

	log.Printf("VPMULHRSW vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMULHRSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
