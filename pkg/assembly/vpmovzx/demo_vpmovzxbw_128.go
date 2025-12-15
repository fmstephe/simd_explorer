package vpmovzx

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovzxbw_128.s
var assemblyVpmovzxbw128 string

//go:embed stub_vpmovzxbw_128.go
var stubVpmovzxbw128 string

type VPMOVZXBW128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVZXBW128() *VPMOVZXBW128 {
	return &VPMOVZXBW128{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPMOVZXBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVZXBW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVZXBW128) Name() string {
	return "VPMOVZXBW (128 bit) "
}

func (v *VPMOVZXBW128) Description() string {
	return "Zero-extend packed 8-bit integers to 16-bit integers."
}

func (v *VPMOVZXBW128) Stub() string {
	return stubVpmovzxbw128
}

func (v *VPMOVZXBW128) Assembly() string {
	return assemblyVpmovzxbw128
}

func (v *VPMOVZXBW128) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [8]uint16{}
	copy(ret[:], number.ToUint16Slice(v.ret.FlatData()))

	vpmovzxbw128(&vals, &ret)

	log.Printf("VPMOVZXBW vals %v ret %v", vals, ret)

	retBytes := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPMOVZXBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
