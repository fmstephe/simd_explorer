package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxwd_128.s
var assemblyVpmovzxwd128 string

//go:embed stub_vpmovzxwd_128.go
var stubVpmovzxwd128 string

type VPMOVZXWD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXWD128() *VPMOVZXWD128 {
	return &VPMOVZXWD128{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMOVZXWD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXWD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXWD128) Name() string {
	return "VPMOVZXWD (128 bit) "
}

func (v *VPMOVZXWD128) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXWD128) Stub() string {
	return stubVpmovzxwd128
}

func (v *VPMOVZXWD128) Assembly() string {
	return assemblyVpmovzxwd128
}

func (v *VPMOVZXWD128) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpmovzxwd128(&vals, &ret)

	log.Printf("VPMOVZXWD vals %v ret %v", vals, ret)

	retBytes := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXWD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
