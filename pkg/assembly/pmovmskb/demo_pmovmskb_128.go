package pmovmskb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmovmskb_128.s
var assemblyPmovmskb128 string

//go:embed stub_pmovmskb_128.go
var stubPmovmskb128 string

type PMOVMSKB128 struct {
}

func (v *PMOVMSKB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 8, 10),
	}
}

func (v *PMOVMSKB128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 16)
}

func (v *PMOVMSKB128) Name() string {
	return "PMOVMSKB (128 bit)"
}

func (v *PMOVMSKB128) Description() string {
	return "Move byte MSBs to mask: packs the most-significant bit of each of 16 bytes into a 16-bit mask (returned as 32-bit value)."
}

func (v *PMOVMSKB128) Stub() string {
	return stubPmovmskb128
}

func (v *PMOVMSKB128) Assembly() string {
	return assemblyPmovmskb128
}

func (v *PMOVMSKB128) Run(inputs [][]byte) (output []byte) {
	vals := [16]uint8{}
	copy(vals[:], inputs[0])

	var ret uint32

	pmovmskb128(&vals, &ret)

	log.Printf("PMOVMSKB128 input %v mask 0x%08x", vals, ret)

	return number.Uint32ToBytes(ret)
}

func (v *PMOVMSKB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
