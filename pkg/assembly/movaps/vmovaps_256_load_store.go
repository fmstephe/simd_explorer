package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_256_loadstore_vmovaps.s
var assembly256LoadStoreVmovaps string

//go:embed stub_256_loadstore_vmovaps.go
var stub256LoadStoreVmovaps string

type VMOVAPS256LoadStore struct {
}

func (v *VMOVAPS256LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMOVAPS256LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMOVAPS256LoadStore) Name() string {
	return "VMOVAPS YMM (256 bit)"
}

func (v *VMOVAPS256LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVAPS256LoadStore) Stub() string {
	return stub256LoadStoreVmovaps
}

func (v *VMOVAPS256LoadStore) Assembly() string {
	return assembly256LoadStoreVmovaps
}

func (v *VMOVAPS256LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	movaps256LoadStoreVmovaps(&floats, &ret)

	log.Printf("VMOVAPS256LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVAPS256LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
