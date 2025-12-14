package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxwq_256.s
var assemblyVpmovzxwq256 string

//go:embed stub_vpmovzxwq_256.go
var stubVpmovzxwq256 string

type VPMOVZXWQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXWQ256() *VPMOVZXWQ256 {
	return &VPMOVZXWQ256{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVZXWQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXWQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXWQ256) Name() string {
	return "VPMOVZXWQ (256 bit) "
}

func (v *VPMOVZXWQ256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXWQ256) Stub() string {
	return stubVpmovzxwq256
}

func (v *VPMOVZXWQ256) Assembly() string {
	return assemblyVpmovzxwq256
}

func (v *VPMOVZXWQ256) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxwq256(&vals, &ret)

	log.Printf("VPMOVZXWQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXWQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
