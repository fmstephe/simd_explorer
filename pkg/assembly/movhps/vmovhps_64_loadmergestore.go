package movhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_64_loadmergestore_vmovhps.s
var assembly64LoadStoreVmovhps string

//go:embed stub_64_loadmergestore_vmovhps.go
var stub64LoadStoreVmovhps string

type VMOVHPS64LoadStore struct {
}

func (v *VMOVHPS64LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *VMOVHPS64LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVHPS64LoadStore) Name() string {
	return "VMOVHPS XMM (2X 64 bit)"
}

func (v *VMOVHPS64LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVHPS64LoadStore) Stub() string {
	return stub64LoadStoreVmovhps
}

func (v *VMOVHPS64LoadStore) Assembly() string {
	return assembly64LoadStoreVmovhps
}

func (v *VMOVHPS64LoadStore) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movhps64LoadMergeStoreVmovhps(&lower, &upper, &ret)

	log.Printf("VMOVHPS64LoadMergeStore input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVHPS64LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
