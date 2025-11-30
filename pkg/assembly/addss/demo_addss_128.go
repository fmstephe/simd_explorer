package addss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_addss_128.s
var assemblyAddss128 string

//go:embed stub_addss_128.go
var stubAddss128 string

type ADDSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewADDSS128() *ADDSS128 {
	return &ADDSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *ADDSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *ADDSS128) Output() *number.Parameter {
	return v.ret
}

func (v *ADDSS128) Name() string {
	return "ADDSS (128 bit) "
}

func (v *ADDSS128) Description() string {
	return "Add scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *ADDSS128) Stub() string {
	return stubAddss128
}

func (v *ADDSS128) Assembly() string {
	return assemblyAddss128
}

func (v *ADDSS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	addss128(&vals1, &vals2, &ret)

	log.Printf("ADDSS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *ADDSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
