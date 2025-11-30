package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_addps_128.s
var assemblyAddps128 string

//go:embed stub_addps_128.go
var stubAddps128 string

type ADDPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewADDPS128() *ADDPS128 {
	return &ADDPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *ADDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *ADDPS128) Output() *number.Parameter {
	return v.ret
}

func (v *ADDPS128) Name() string {
	return "ADDPS XMM (128 bit)"
}

func (v *ADDPS128) Description() string {
	return "Add packed single-precision floats in XMM, lane-wise."
}

func (v *ADDPS128) Stub() string {
	return stubAddps128
}

func (v *ADDPS128) Assembly() string {
	return assemblyAddps128
}

func (v *ADDPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	addps128(&vals1, &vals2, &ret)

	log.Printf("ADDPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *ADDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
