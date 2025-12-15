package vpmuldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmuldq_128.s
var assemblyVpmuldq128 string

//go:embed stub_vpmuldq_128.go
var stubVpmuldq128 string

type VPMULDQ128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULDQ128() *VPMULDQ128 {
	return &VPMULDQ128{
		vals1: number.NewNamedIntParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 64, 10),
	}
}

func (v *VPMULDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULDQ128) Name() string {
	return "VPMULDQ (128 bit) "
}

func (v *VPMULDQ128) Description() string {
	return "Multiply pairs of signed 32-bit integers to 64-bit results (even-indexed pairs). ret[i] = vals1[2*i] * vals2[2*i]."
}

func (v *VPMULDQ128) Stub() string {
	return stubVpmuldq128
}

func (v *VPMULDQ128) Assembly() string {
	return assemblyVpmuldq128
}

func (v *VPMULDQ128) Run() {
	vals1 := [4]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [4]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))
	ret := [2]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmuldq128(&vals1, &vals2, &ret)

	log.Printf("VPMULDQ vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMULDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
