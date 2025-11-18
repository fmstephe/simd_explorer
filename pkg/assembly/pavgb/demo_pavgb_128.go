package pavgb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pavgb_128.s
var assemblyPavgb128 string

//go:embed stub_pavgb_128.go
var stubPavgb128 string

type PAVGB128 struct {
}

func (v *PAVGB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *PAVGB128) Output() *number.Parameter {
	return number.NewUintParameter(128, 8, 10)
}

func (v *PAVGB128) Name() string {
	return "PAVGB (128 bit)"
}

func (v *PAVGB128) Description() string {
	return "Average of packed unsigned bytes with rounding: (a+b+1)>>1."
}

func (v *PAVGB128) Stub() string {
	return stubPavgb128
}

func (v *PAVGB128) Assembly() string {
	return assemblyPavgb128
}

func (v *PAVGB128) Run(inputs [][]byte) (output []byte) {
	b1 := [16]uint8{}
	copy(b1[:], inputs[0])
	b2 := [16]uint8{}
	copy(b2[:], inputs[1])

	ret := [16]uint8{}

	pavgb128(&b1, &b2, &ret)

	log.Printf("PAVGB128 input %v %v output %v", b1, b2, ret)

	return ret[:]
}

func (v *PAVGB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
