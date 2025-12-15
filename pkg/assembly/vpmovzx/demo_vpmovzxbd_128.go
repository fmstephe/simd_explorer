package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbd_128.s
var assemblyVpmovzxbd128 string

//go:embed stub_vpmovzxbd_128.go
var stubVpmovzxbd128 string

type VPMOVZXBD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBD128() *VPMOVZXBD128 {
	return &VPMOVZXBD128{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMOVZXBD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBD128) Name() string {
	return "VPMOVZXBD (128 bit) "
}

func (v *VPMOVZXBD128) Description() string {
	return "Zero-extend packed 8-bit integers to 32-bit integers."
}

func (v *VPMOVZXBD128) Stub() string {
	return stubVpmovzxbd128
}

func (v *VPMOVZXBD128) Assembly() string {
	return assemblyVpmovzxbd128
}

func (v *VPMOVZXBD128) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [4]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vpmovzxbd128(&vals, &ret)

	log.Printf("VPMOVZXBD vals %v ret %v", vals, ret)

	retBytes := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
