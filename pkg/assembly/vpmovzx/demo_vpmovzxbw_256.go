package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbw_256.s
var assemblyVpmovzxbw256 string

//go:embed stub_vpmovzxbw_256.go
var stubVpmovzxbw256 string

type VPMOVZXBW256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBW256() *VPMOVZXBW256 {
	return &VPMOVZXBW256{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPMOVZXBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBW256) Name() string {
	return "VPMOVZXBW (256 bit) "
}

func (v *VPMOVZXBW256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXBW256) Stub() string {
	return stubVpmovzxbw256
}

func (v *VPMOVZXBW256) Assembly() string {
	return assemblyVpmovzxbw256
}

func (v *VPMOVZXBW256) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint16{}
	copy(ret[:], number.ToUint16Slice(v.ret.FlatData()))

	vpmovzxbw256(&vals, &ret)

	log.Printf("VPMOVZXBW vals %v ret %v", vals, ret)

	retBytes := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
