package movhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_64_loadmergestore_movhps.s
var assembly64LoadStoreMovhps string

//go:embed stub_64_loadmergestore_movhps.go
var stub64LoadStoreMovhps string

type MOVHPS64LoadStore struct {
}

func (v *MOVHPS64LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVHPS64LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVHPS64LoadStore) Name() string {
	return "MOVHPS XMM (2X 64 bit)"
}

func (v *MOVHPS64LoadStore) Description() string {
	return "TODO"
}

func (v *MOVHPS64LoadStore) Stub() string {
	return stub64LoadStoreMovhps
}

func (v *MOVHPS64LoadStore) Assembly() string {
	return assembly64LoadStoreMovhps
}

func (v *MOVHPS64LoadStore) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movhps64LoadMergeStoreMovhps(&lower, &upper, &ret)

	log.Printf("MOVHPS64LoadMergeStore input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVHPS64LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
