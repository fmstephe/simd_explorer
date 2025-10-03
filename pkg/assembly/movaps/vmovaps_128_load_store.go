package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_vmovaps.s
var assembly128LoadStoreVmovaps string

//go:embed stub_128_loadstore_vmovaps.go
var stub128LoadStoreVmovaps string

type VMOVAPS128LoadStore struct {
}

func (v *VMOVAPS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVAPS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVAPS128LoadStore) Name() string {
	return "VMOVAPS XMM (128 bit)"
}

func (v *VMOVAPS128LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVAPS128LoadStore) Stub() string {
	return stub128LoadStoreVmovaps
}

func (v *VMOVAPS128LoadStore) Assembly() string {
	return assembly128LoadStoreVmovaps
}

func (v *VMOVAPS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movaps128LoadStoreVmovaps(&floats, &ret)

	log.Printf("VMOVAPS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVAPS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
