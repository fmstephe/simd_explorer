package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_64_loadmergestore_movlps.s
var assembly64LoadStoreMovlps string

//go:embed stub_64_loadmergestore_movlps.go
var stub64LoadStoreMovlps string

type MOVLPS64LoadStore struct {
}

func (v *MOVLPS64LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVLPS64LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVLPS64LoadStore) Name() string {
	return "MOVLPS XMM (2X 64 bit)"
}

func (v *MOVLPS64LoadStore) Description() string {
	return "TODO"
}

func (v *MOVLPS64LoadStore) Stub() string {
	return stub64LoadStoreMovlps
}

func (v *MOVLPS64LoadStore) Assembly() string {
	return assembly64LoadStoreMovlps
}

func (v *MOVLPS64LoadStore) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movlps64LoadMergeStoreMovlps(&lower, &upper, &ret)

	log.Printf("MOVLPS64LoadMergeStore input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVLPS64LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
