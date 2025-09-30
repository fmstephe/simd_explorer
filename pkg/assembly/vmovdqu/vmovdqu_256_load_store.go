package vmovdqu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_256_loadstore.s
var assembly256 string

//go:embed stub_256_loadstore.go
var stub256 string

type VMOVDQU256LoadStore struct {
}

func (v *VMOVDQU256LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10),
	}
}

func (v *VMOVDQU256LoadStore) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 10)
}

func (v *VMOVDQU256LoadStore) Name() string {
	return "VMOVDQU YMM (256 bit)"
}

func (v *VMOVDQU256LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVDQU256LoadStore) Stub() string {
	return stub256
}

func (v *VMOVDQU256LoadStore) Assembly() string {
	return assembly256
}

func (v *VMOVDQU256LoadStore) Run(inputs [][]byte) (output []byte) {
	uints := [8]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [8]uint32{}

	vmovdqu256LoadStore(&uints, &ret)

	log.Printf("VMOVDQU256LoadStore input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQU256LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
