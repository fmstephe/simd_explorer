package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_vmovups.s
var assembly128LoadStoreVmovups string

//go:embed stub_128_loadstore_vmovups.go
var stub128LoadStoreVmovups string

type VMOVUPS128LoadStore struct {
}

func (v *VMOVUPS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVUPS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVUPS128LoadStore) Name() string {
	return "VMOVUPS XMM (128 bit)"
}

func (v *VMOVUPS128LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVUPS128LoadStore) Stub() string {
	return stub128LoadStoreVmovups
}

func (v *VMOVUPS128LoadStore) Assembly() string {
	return assembly128LoadStoreVmovups
}

func (v *VMOVUPS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movups128LoadStoreVmovups(&floats, &ret)

	log.Printf("VMOVUPS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVUPS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
