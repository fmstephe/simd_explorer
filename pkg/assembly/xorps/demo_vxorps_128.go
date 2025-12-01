package xorps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vxorps_128.s
var assemblyVxorps128 string

//go:embed stub_vxorps_128.go
var stubVxorps128 string

type VXORPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVXORPS128() *VXORPS128 {
	return &VXORPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VXORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VXORPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VXORPS128) Name() string {
	return "VXORPS (128 bit)"
}

func (v *VXORPS128) Description() string {
	return "Bitwise XOR of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VXORPS128) Stub() string {
	return stubVxorps128
}

func (v *VXORPS128) Assembly() string {
	return assemblyVxorps128
}

func (v *VXORPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [4]float32{}

	vxorps128(&vals1, &vals2, &ret)

	log.Printf("VXORPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VXORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
