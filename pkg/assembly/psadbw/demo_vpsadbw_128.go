package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsadbw_128.s
var assemblyVpsadbw128 string

//go:embed stub_vpsadbw_128.go
var stubVpsadbw128 string

type VPSADBW128 struct {
}

func (v *VPSADBW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *VPSADBW128) Output() *number.Parameter {
	return number.NewUintParameter(128, 16, 10)
}

func (v *VPSADBW128) Name() string {
	return "VPSADBW (128 bit)"
}

func (v *VPSADBW128) Description() string {
	return "Packed sum of absolute byte differences (VEX); two 64-bit lane sums."
}

func (v *VPSADBW128) Stub() string {
	return stubVpsadbw128
}

func (v *VPSADBW128) Assembly() string {
	return assemblyVpsadbw128
}

func (v *VPSADBW128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [8]uint16{}

	vpsadbw128(&b1, &b2, &ret)

	log.Printf("VPSADBW128 input %v %v output %v", b1, b2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPSADBW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
