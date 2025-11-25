package vextracti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vextracti128_256_zero.s
var assemblyVextracti128256Zero string

//go:embed stub_vextracti128_256_zero.go
var stubVextracti128256Zero string

type VEXTRACTI128256ZERO struct {
}

func (v *VEXTRACTI128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10), // vals256 (8x u32)
	}
}

func (v *VEXTRACTI128256ZERO) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VEXTRACTI128256ZERO) Name() string {
	return "VEXTRACTI128 (256 bit) zero"
}

func (v *VEXTRACTI128256ZERO) Description() string {
	return "Extract lower 128-bit lane (0) from YMM to memory."
}

func (v *VEXTRACTI128256ZERO) Stub() string {
	return stubVextracti128256Zero
}

func (v *VEXTRACTI128256ZERO) Assembly() string {
	return assemblyVextracti128256Zero
}

func (v *VEXTRACTI128256ZERO) Run(inputs [][]byte) (output []byte) {
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(inputs[0]))
	var ret [4]uint32
	vextracti128256Zero(&vals256, &ret)
	log.Printf("VEXTRACTI128256ZERO vals256 %v output %v", vals256, ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VEXTRACTI128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
