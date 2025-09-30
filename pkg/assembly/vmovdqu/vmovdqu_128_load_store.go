package vmovdqu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore.s
var assembly128 string

//go:embed stub_128_loadstore.go
var stub128 string

type VMOVDQU128LoadStore struct {
}

func (v *VMOVDQU128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 10),
	}
}

func (v *VMOVDQU128LoadStore) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VMOVDQU128LoadStore) Name() string {
	return "VMOVDQU XMM (128 bit)"
}

func (v *VMOVDQU128LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVDQU128LoadStore) Stub() string {
	return stub128
}

func (v *VMOVDQU128LoadStore) Assembly() string {
	return assembly128
}

func (v *VMOVDQU128LoadStore) Run(inputs [][]byte) (output []byte) {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [4]uint32{}

	vmovdqu128LoadStore(&uints, &ret)

	log.Printf("VMOVDQU128LoadStore input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQU128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
