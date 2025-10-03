package movss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_vmovss.s
var assembly128LoadStoreVmovss string

//go:embed stub_128_loadstore_vmovss.go
var stub128LoadStoreVmovss string

type VMOVSS128LoadStore struct {
}

func (v *VMOVSS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVSS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVSS128LoadStore) Name() string {
	return "VMOVSS XMM (128 bit)"
}

func (v *VMOVSS128LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVSS128LoadStore) Stub() string {
	return stub128LoadStoreVmovss
}

func (v *VMOVSS128LoadStore) Assembly() string {
	return assembly128LoadStoreVmovss
}

func (v *VMOVSS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movss128LoadStoreVmovss(&floats, &ret)

	log.Printf("VMOVSS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVSS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
