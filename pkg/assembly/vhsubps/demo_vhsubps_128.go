package vhsubps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhsubps_128.s
var assemblyVhsubps128 string

//go:embed stub_vhsubps_128.go
var stubVhsubps128 string

type VHSUBPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHSUBPS128() *VHSUBPS128 {
	return &VHSUBPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VHSUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHSUBPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VHSUBPS128) Name() string {
	return "VHSUBPS (128 bit) "
}

func (v *VHSUBPS128) Description() string {
	return "Horizontal subtract of packed single-precision floats. " +
		"Computes pairwise differences within each source (a0-a1, a2-a3) and (b0-b1, b2-b3), writes results to ret."
}

func (v *VHSUBPS128) Stub() string {
	return stubVhsubps128
}

func (v *VHSUBPS128) Assembly() string {
	return assemblyVhsubps128
}

func (v *VHSUBPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vhsubps128(&vals1, &vals2, &ret)

	log.Printf("VHSUBPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHSUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
