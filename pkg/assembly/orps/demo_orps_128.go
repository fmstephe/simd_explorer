package orps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_orps_128.s
var assemblyOrps128 string

//go:embed stub_orps_128.go
var stubOrps128 string

type ORPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewORPS128() *ORPS128 {
	return &ORPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *ORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *ORPS128) Output() *number.Parameter {
	return v.ret
}

func (v *ORPS128) Name() string {
	return "ORPS (128 bit)"
}

func (v *ORPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ORPS128) Stub() string {
	return stubOrps128
}

func (v *ORPS128) Assembly() string {
	return assemblyOrps128
}

func (v *ORPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	orps128(&vals1, &vals2, &ret)

	log.Printf("ORPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *ORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
