package xorps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vxorps_256.s
var assemblyVxorps256 string

//go:embed stub_vxorps_256.go
var stubVxorps256 string

type VXORPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVXORPS256() *VXORPS256 {
	return &VXORPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VXORPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VXORPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VXORPS256) Name() string {
	return "VXORPS (256 bit)"
}

func (v *VXORPS256) Description() string {
	return "Bitwise XOR of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VXORPS256) Stub() string {
	return stubVxorps256
}

func (v *VXORPS256) Assembly() string {
	return assemblyVxorps256
}

func (v *VXORPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [8]float32{}

	vxorps256(&vals1, &vals2, &ret)

	log.Printf("VXORPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VXORPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
