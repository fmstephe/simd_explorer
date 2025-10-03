package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_256_loadstore_vmovups.s
var assembly256LoadStoreVmovups string

//go:embed stub_256_loadstore_vmovups.go
var stub256LoadStoreVmovups string

type VMOVUPS256LoadStore struct {
}

func (v *VMOVUPS256LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMOVUPS256LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMOVUPS256LoadStore) Name() string {
	return "VMOVUPS YMM (256 bit)"
}

func (v *VMOVUPS256LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVUPS256LoadStore) Stub() string {
	return stub256LoadStoreVmovups
}

func (v *VMOVUPS256LoadStore) Assembly() string {
	return assembly256LoadStoreVmovups
}

func (v *VMOVUPS256LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	movups256LoadStoreVmovups(&floats, &ret)

	log.Printf("VMOVUPS256LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVUPS256LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
