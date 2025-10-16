package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_64_loadmergestore_vmovlps.s
var assembly64LoadStoreVmovlps string

//go:embed stub_64_loadmergestore_vmovlps.go
var stub64LoadStoreVmovlps string

type VMOVLPS64LoadStore struct {
}

func (v *VMOVLPS64LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *VMOVLPS64LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVLPS64LoadStore) Name() string {
	return "VMOVLPS XMM (2X 64 bit)"
}

func (v *VMOVLPS64LoadStore) Description() string {
	return "TODO"
}

func (v *VMOVLPS64LoadStore) Stub() string {
	return stub64LoadStoreVmovlps
}

func (v *VMOVLPS64LoadStore) Assembly() string {
	return assembly64LoadStoreVmovlps
}

func (v *VMOVLPS64LoadStore) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	movlps64LoadMergeStoreVmovlps(&lower, &upper, &ret)

	log.Printf("VMOVLPS64LoadMergeStore input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVLPS64LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
