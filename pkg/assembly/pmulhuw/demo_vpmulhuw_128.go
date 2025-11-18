package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmulhuw_128.s
var assemblyVpmulhuw128 string

//go:embed stub_vpmulhuw_128.go
var stubVpmulhuw128 string

type VPMULHUW128 struct {
}

func (v *VPMULHUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *VPMULHUW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *VPMULHUW128) Name() string {
	return "VPMULHUW (128 bit)"
}

func (v *VPMULHUW128) Description() string {
	return "Packed multiply unsigned 16-bit integers (VEX); keep the high 16 bits of each 32-bit product."
}

func (v *VPMULHUW128) Stub() string {
	return stubVpmulhuw128
}

func (v *VPMULHUW128) Assembly() string {
	return assemblyVpmulhuw128
}

func (v *VPMULHUW128) Run(inputs [][]byte) (output []byte) {
	u1 := [8]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [8]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [8]uint16{}

	vpmulhuw128(&u1, &u2, &ret)

	log.Printf("VPMULHUW128 input %v %v output %v", u1, u2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPMULHUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
