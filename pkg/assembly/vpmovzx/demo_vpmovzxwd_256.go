package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxwd_256.s
var assemblyVpmovzxwd256 string

//go:embed stub_vpmovzxwd_256.go
var stubVpmovzxwd256 string

type VPMOVZXWD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXWD256() *VPMOVZXWD256 {
	return &VPMOVZXWD256{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMOVZXWD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXWD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXWD256) Name() string {
	return "VPMOVZXWD (256 bit) "
}

func (v *VPMOVZXWD256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXWD256) Stub() string {
	return stubVpmovzxwd256
}

func (v *VPMOVZXWD256) Assembly() string {
	return assemblyVpmovzxwd256
}

func (v *VPMOVZXWD256) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))
	ret := [8]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpmovzxwd256(&vals, &ret)

	log.Printf("VPMOVZXWD vals %v ret %v", vals, ret)

	retBytes := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXWD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
