package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhuw_256.s
var assemblyVpmulhuw256 string

//go:embed stub_vpmulhuw_256.go
var stubVpmulhuw256 string

type VPMULHUW256 struct {
}

func (v *VPMULHUW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 16, 10),
		number.NewUintParameter(256, 16, 10),
	}
}

func (v *VPMULHUW256) Output() *number.Parameter {
	return number.NewUintParameter(256, 16, 10)
}

func (v *VPMULHUW256) Name() string {
	return "VPMULHUW (256 bit)"
}

func (v *VPMULHUW256) Description() string {
	return "Packed multiply unsigned 16-bit integers (VEX, per 128-bit lane); keep the high 16 bits."
}

func (v *VPMULHUW256) Stub() string {
	return stubVpmulhuw256
}

func (v *VPMULHUW256) Assembly() string {
	return assemblyVpmulhuw256
}

func (v *VPMULHUW256) Run(inputs [][]byte) (output []byte) {
	u1 := [16]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [16]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [16]uint16{}

	vpmulhuw256(&u1, &u2, &ret)

	log.Printf("VPMULHUW256 input %v %v output %v", u1, u2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPMULHUW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
