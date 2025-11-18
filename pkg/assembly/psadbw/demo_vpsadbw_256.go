package psadbw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsadbw_256.s
var assemblyVpsadbw256 string

//go:embed stub_vpsadbw_256.go
var stubVpsadbw256 string

type VPSADBW256 struct {
}

func (v *VPSADBW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 8, 10),
		number.NewUintParameter(256, 8, 10),
	}
}

func (v *VPSADBW256) Output() *number.Parameter {
	return number.NewUintParameter(256, 16, 10)
}

func (v *VPSADBW256) Name() string {
	return "VPSADBW (256 bit)"
}

func (v *VPSADBW256) Description() string {
	return "Packed sum of absolute byte differences (VEX); four 64-bit lane sums."
}

func (v *VPSADBW256) Stub() string {
	return stubVpsadbw256
}

func (v *VPSADBW256) Assembly() string {
	return assemblyVpsadbw256
}

func (v *VPSADBW256) Run(inputs [][]byte) (output []byte) {
	b1 := [32]uint8{}
	copy(b1[:], inputs[0])
	b2 := [32]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint16{}

	vpsadbw256(&b1, &b2, &ret)

	log.Printf("VPSADBW256 input %v %v output %v", b1, b2, ret)

	bytes := []byte{}
	for _, v := range ret {
		bytes = append(bytes, number.Uint16ToBytes(v)...)
	}
	return bytes
}

func (v *VPSADBW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
