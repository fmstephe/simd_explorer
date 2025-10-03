package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_movups.s
var assembly128LoadStoreMovups string

//go:embed stub_128_loadstore_movups.go
var stub128LoadStoreMovups string

type MOVUPS128LoadStore struct {
}

func (v *MOVUPS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVUPS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVUPS128LoadStore) Name() string {
	return "MOVUPS YMM (128 bit)"
}

func (v *MOVUPS128LoadStore) Description() string {
	return "TODO"
}

func (v *MOVUPS128LoadStore) Stub() string {
	return stub128LoadStoreMovups
}

func (v *MOVUPS128LoadStore) Assembly() string {
	return assembly128LoadStoreMovups
}

func (v *MOVUPS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movups128LoadStoreMovups(&floats, &ret)

	log.Printf("MOVUPS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVUPS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
