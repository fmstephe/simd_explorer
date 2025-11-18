package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_psadbw_128.s
var assemblyPsadbw128 string

//go:embed stub_psadbw_128.go
var stubPsadbw128 string

type PSADBW128 struct {
}

func (v *PSADBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *PSADBW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *PSADBW128) Name() string {
	return "PSADBW (128 bit)"
}

func (v *PSADBW128) Description() string {
	return "Packed sum of absolute byte differences; two 64-bit lane sums."
}

func (v *PSADBW128) Stub() string {
	return stubPsadbw128
}

func (v *PSADBW128) Assembly() string {
	return assemblyPsadbw128
}

func (v *PSADBW128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [8]uint16{}

	psadbw128(&b1, &b2, &ret)

	log.Printf("PSADBW128 input %v %v output %v", b1, b2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *PSADBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
