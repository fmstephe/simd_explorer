package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxdq_256.s
var assemblyVpmovzxdq256 string

//go:embed stub_vpmovzxdq_256.go
var stubVpmovzxdq256 string

type VPMOVZXDQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXDQ256() *VPMOVZXDQ256 {
	return &VPMOVZXDQ256{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVZXDQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXDQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXDQ256) Name() string {
	return "VPMOVZXDQ (256 bit) "
}

func (v *VPMOVZXDQ256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXDQ256) Stub() string {
	return stubVpmovzxdq256
}

func (v *VPMOVZXDQ256) Assembly() string {
	return assemblyVpmovzxdq256
}

func (v *VPMOVZXDQ256) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxdq256(&vals, &ret)

	log.Printf("VPMOVZXDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXDQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
