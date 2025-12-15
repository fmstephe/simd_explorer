package vpmuldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmuldq_256.s
var assemblyVpmuldq256 string

//go:embed stub_vpmuldq_256.go
var stubVpmuldq256 string

type VPMULDQ256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMULDQ256() *VPMULDQ256 {
	return &VPMULDQ256{
		vals1: number.NewNamedIntParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 64, 10),
	}
}

func (v *VPMULDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMULDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMULDQ256) Name() string {
	return "VPMULDQ (256 bit) "
}

func (v *VPMULDQ256) Description() string {
	return "Multiply pairs of signed 32-bit integers to 64-bit results (even-indexed pairs), per 128-bit lane. ret[i] = vals1[2*i] * vals2[2*i]."
}

func (v *VPMULDQ256) Stub() string {
	return stubVpmuldq256
}

func (v *VPMULDQ256) Assembly() string {
	return assemblyVpmuldq256
}

func (v *VPMULDQ256) Run() {
	vals1 := [8]int32{}
	copy(vals1[:], number.ToInt32Slice(v.vals1.FlatData()))
	vals2 := [8]int32{}
	copy(vals2[:], number.ToInt32Slice(v.vals2.FlatData()))
	ret := [4]int64{}
	copy(ret[:], number.ToInt64Slice(v.ret.FlatData()))

	vpmuldq256(&vals1, &vals2, &ret)

	log.Printf("VPMULDQ vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Int64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMULDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
