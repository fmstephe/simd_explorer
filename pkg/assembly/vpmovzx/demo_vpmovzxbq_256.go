package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbq_256.s
var assemblyVpmovzxbq256 string

//go:embed stub_vpmovzxbq_256.go
var stubVpmovzxbq256 string

type VPMOVZXBQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBQ256() *VPMOVZXBQ256 {
	return &VPMOVZXBQ256{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPMOVZXBQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBQ256) Name() string {
	return "VPMOVZXBQ (256 bit) "
}

func (v *VPMOVZXBQ256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXBQ256) Stub() string {
	return stubVpmovzxbq256
}

func (v *VPMOVZXBQ256) Assembly() string {
	return assemblyVpmovzxbq256
}

func (v *VPMOVZXBQ256) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [4]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxbq256(&vals, &ret)

	log.Printf("VPMOVZXBQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
