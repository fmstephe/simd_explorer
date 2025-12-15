package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbq_128.s
var assemblyVpmovzxbq128 string

//go:embed stub_vpmovzxbq_128.go
var stubVpmovzxbq128 string

type VPMOVZXBQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBQ128() *VPMOVZXBQ128 {
	return &VPMOVZXBQ128{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMOVZXBQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBQ128) Name() string {
	return "VPMOVZXBQ (128 bit) "
}

func (v *VPMOVZXBQ128) Description() string {
	return "Zero-extend packed 8-bit integers to 64-bit integers."
}

func (v *VPMOVZXBQ128) Stub() string {
	return stubVpmovzxbq128
}

func (v *VPMOVZXBQ128) Assembly() string {
	return assemblyVpmovzxbq128
}

func (v *VPMOVZXBQ128) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [2]uint64{}
	copy(ret[:], number.ToUint64Slice(v.ret.FlatData()))

	vpmovzxbq128(&vals, &ret)

	log.Printf("VPMOVZXBQ vals %v ret %v", vals, ret)

	retBytes := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
