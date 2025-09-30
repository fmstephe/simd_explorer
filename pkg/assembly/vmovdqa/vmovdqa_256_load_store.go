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

type VMOVDQA256LoadStore struct {
}

func (v *VMOVDQA256LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10),
	}
}

func (v *VMOVDQA256LoadStore) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 10)
}

func (v *VMOVDQA256LoadStore) Name() string {
	return "VMOVDQA YMM (256 bit)"
}

func (v *VMOVDQA256LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVDQA256LoadStore) Stub() string {
	return stub256
}

func (v *VMOVDQA256LoadStore) Assembly() string {
	return assembly256
}

func (v *VMOVDQA256LoadStore) Run(inputs [][]byte) (output []byte) {
	uints := [8]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [8]uint32{}

	vmovdqu256LoadStore(&uints, &ret)

	log.Printf("VMOVDQA256LoadStore input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQA256LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
