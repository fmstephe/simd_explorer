package vhaddps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhaddps_128.s
var assemblyVhaddps128 string

//go:embed stub_vhaddps_128.go
var stubVhaddps128 string

type VHADDPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHADDPS128() *VHADDPS128 {
	return &VHADDPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VHADDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHADDPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VHADDPS128) Name() string {
	return "VHADDPS (128 bit) "
}

func (v *VHADDPS128) Description() string {
	return "Horizontal add of packed single-precision floats. " +
		"Computes pairwise sums within each source (a0+a1, a2+a3) and (b0+b1, b2+b3), writes results to ret."
}

func (v *VHADDPS128) Stub() string {
	return stubVhaddps128
}

func (v *VHADDPS128) Assembly() string {
	return assemblyVhaddps128
}

func (v *VHADDPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vhaddps128(&vals1, &vals2, &ret)

	log.Printf("VHADDPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHADDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
