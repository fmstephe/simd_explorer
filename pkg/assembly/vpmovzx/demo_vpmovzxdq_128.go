package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxdq_128.s
var assemblyVpmovzxdq128 string

//go:embed stub_vpmovzxdq_128.go
var stubVpmovzxdq128 string

type VPMOVZXDQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXDQ128() *VPMOVZXDQ128 {
	return &VPMOVZXDQ128{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVZXDQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXDQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXDQ128) Name() string {
	return "VPMOVZXDQ (128 bit) "
}

func (v *VPMOVZXDQ128) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXDQ128) Stub() string {
	return stubVpmovzxdq128
}

func (v *VPMOVZXDQ128) Assembly() string {
	return assemblyVpmovzxdq128
}

func (v *VPMOVZXDQ128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxdq128(&vals, &ret)

	log.Printf("VPMOVZXDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXDQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
