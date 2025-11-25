package vextracti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vextracti128_256_one.s
var assemblyVextracti128256One string

//go:embed stub_vextracti128_256_one.go
var stubVextracti128256One string

type VEXTRACTI128256ONE struct {
}

func (v *VEXTRACTI128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10), // vals256 (8x u32)
	}
}

func (v *VEXTRACTI128256ONE) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VEXTRACTI128256ONE) Name() string {
	return "VEXTRACTI128 (256 bit) one"
}

func (v *VEXTRACTI128256ONE) Description() string {
	return "Extract upper 128-bit lane (1) from YMM to memory."
}

func (v *VEXTRACTI128256ONE) Stub() string {
	return stubVextracti128256One
}

func (v *VEXTRACTI128256ONE) Assembly() string {
	return assemblyVextracti128256One
}

func (v *VEXTRACTI128256ONE) Run(inputs [][]byte) (output []byte) {
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(inputs[0]))
	var ret [4]uint32
	vextracti128256One(&vals256, &ret)
	log.Printf("VEXTRACTI128256ONE vals256 %v output %v", vals256, ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VEXTRACTI128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
