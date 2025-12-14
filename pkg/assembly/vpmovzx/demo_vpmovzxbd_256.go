package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbd_256.s
var assemblyVpmovzxbd256 string

//go:embed stub_vpmovzxbd_256.go
var stubVpmovzxbd256 string

type VPMOVZXBD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBD256() *VPMOVZXBD256 {
	return &VPMOVZXBD256{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMOVZXBD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBD256) Name() string {
	return "VPMOVZXBD (256 bit) "
}

func (v *VPMOVZXBD256) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *VPMOVZXBD256) Stub() string {
	return stubVpmovzxbd256
}

func (v *VPMOVZXBD256) Assembly() string {
	return assemblyVpmovzxbd256
}

func (v *VPMOVZXBD256) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [8]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpmovzxbd256(&vals, &ret)

	log.Printf("VPMOVZXBD vals %v ret %v", vals, ret)

	retBytes := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
