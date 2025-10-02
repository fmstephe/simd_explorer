package vmovdqa

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

type VMOVDQA128LoadStore struct {
}

func (v *VMOVDQA128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 10),
	}
}

func (v *VMOVDQA128LoadStore) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VMOVDQA128LoadStore) Name() string {
	return "VMOVDQA XMM (128 bit)"
}

func (v *VMOVDQA128LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVDQA128LoadStore) Stub() string {
	return stub128
}

func (v *VMOVDQA128LoadStore) Assembly() string {
	return assembly128
}

func (v *VMOVDQA128LoadStore) Run(inputs [][]byte) (output []byte) {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [4]uint32{}

	vmovdqa128LoadStore(&uints, &ret)

	log.Printf("VMOVDQA128LoadStore input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQA128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
