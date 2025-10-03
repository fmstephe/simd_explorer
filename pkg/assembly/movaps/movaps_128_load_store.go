package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_movaps.s
var assembly128LoadStoreMovaps string

//go:embed stub_128_loadstore_movaps.go
var stub128LoadStoreMovaps string

type MOVAPS128LoadStore struct {
}

func (v *MOVAPS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVAPS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVAPS128LoadStore) Name() string {
	return "MOVAPS YMM (128 bit)"
}

func (v *MOVAPS128LoadStore) Description() string {
	return "TODO"
}

func (v *MOVAPS128LoadStore) Stub() string {
	return stub128LoadStoreMovaps
}

func (v *MOVAPS128LoadStore) Assembly() string {
	return assembly128LoadStoreMovaps
}

func (v *MOVAPS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movaps128LoadStoreMovaps(&floats, &ret)

	log.Printf("MOVAPS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVAPS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
