package vpmulhrsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhrsw_256.s
var assemblyVpmulhrsw256 string

//go:embed stub_vpmulhrsw_256.go
var stubVpmulhrsw256 string

type VPMULHRSW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULHRSW256() *VPMULHRSW256 {
	return &VPMULHRSW256{
		vals1: number.NewNamedIntParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 16, 10),
	}
}

func (v *VPMULHRSW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULHRSW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULHRSW256) Name() string {
	return "VPMULHRSW (256 bit) "
}

func (v *VPMULHRSW256) Description() string {
	return "Multiply high with round and scale (signed 16→16). " +
		"For each lane, compute (a*b + 0x8000) >> 16 (signed) and store 16-bit results in ret; operates per 128-bit lane."
}

func (v *VPMULHRSW256) Stub() string {
	return stubVpmulhrsw256
}

func (v *VPMULHRSW256) Assembly() string {
	return assemblyVpmulhrsw256
}

func (v *VPMULHRSW256) Run() {
	vals1 := [16]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [16]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))
	ret := [16]int16{}
	copy(ret[:], number.ToInt16Slice(v.ret.FlatData()))

	vpmulhrsw256(&vals1, &vals2, &ret)

	log.Printf("VPMULHRSW vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMULHRSW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
