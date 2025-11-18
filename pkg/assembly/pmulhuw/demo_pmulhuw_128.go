package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmulhuw_128.s
var assemblyPmulhuw128 string

//go:embed stub_pmulhuw_128.go
var stubPmulhuw128 string

type PMULHUW128 struct {
}

func (v *PMULHUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 16, 10),
		number.NewUintParameter(128, 16, 10),
	}
}

func (v *PMULHUW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PMULHUW128) Name() string {
	return "PMULHUW (128 bit)"
}

func (v *PMULHUW128) Description() string {
	return "Packed multiply unsigned 16-bit integers; keep the high 16 bits of each 32-bit product."
}

func (v *PMULHUW128) Stub() string {
	return stubPmulhuw128
}

func (v *PMULHUW128) Assembly() string {
	return assemblyPmulhuw128
}

func (v *PMULHUW128) Run(inputs [][]byte) (output []byte) {
	u1 := [8]uint16{}
	copy(u1[:], number.ToUint16Slice(inputs[0]))
	u2 := [8]uint16{}
	copy(u2[:], number.ToUint16Slice(inputs[1]))

	ret := [8]uint16{}

	pmulhuw128(&u1, &u2, &ret)

	log.Printf("PMULHUW128 input %v %v output %v", u1, u2, ret)

	// Return as uint16 bytes (decimal in UI)
	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PMULHUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
