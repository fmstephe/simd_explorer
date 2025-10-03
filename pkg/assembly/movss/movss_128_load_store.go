package movss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_loadstore_movss.s
var assembly128LoadStoreMovss string

//go:embed stub_128_loadstore_movss.go
var stub128LoadStoreMovss string

type MOVSS128LoadStore struct {
}

func (v *MOVSS128LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVSS128LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVSS128LoadStore) Name() string {
	return "MOVSS YMM (128 bit)"
}

func (v *MOVSS128LoadStore) Description() string {
	return "TODO"
}

func (v *MOVSS128LoadStore) Stub() string {
	return stub128LoadStoreMovss
}

func (v *MOVSS128LoadStore) Assembly() string {
	return assembly128LoadStoreMovss
}

func (v *MOVSS128LoadStore) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movss128LoadStoreMovss(&floats, &ret)

	log.Printf("MOVSS128LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVSS128LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
