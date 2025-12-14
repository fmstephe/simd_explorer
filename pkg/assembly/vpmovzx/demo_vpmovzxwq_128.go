package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxwq_128.s
var assemblyVpmovzxwq128 string

//go:embed stub_vpmovzxwq_128.go
var stubVpmovzxwq128 string

type VPMOVZXWQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXWQ128() *VPMOVZXWQ128 {
	return &VPMOVZXWQ128{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVZXWQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXWQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXWQ128) Name() string {
	return "VPMOVZXWQ (128 bit) "
}

func (v *VPMOVZXWQ128) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXWQ128) Stub() string {
	return stubVpmovzxwq128
}

func (v *VPMOVZXWQ128) Assembly() string {
	return assemblyVpmovzxwq128
}

func (v *VPMOVZXWQ128) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxwq128(&vals, &ret)

	log.Printf("VPMOVZXWQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXWQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
