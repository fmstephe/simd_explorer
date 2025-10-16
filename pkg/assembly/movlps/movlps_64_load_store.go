package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_64_loadstore_movlps.s
var assembly64LoadStoreMovlps string

//go:embed stub_64_loadstore_movlps.go
var stub64LoadStoreMovlps string

type MOVLPS64LoadStore struct {
}

func (v *MOVLPS64LoadStore) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
	}
}

func (v *MOVLPS64LoadStore) Output() *number.Parameter {
	return number.NewFloatParameter(64, 32)
}

func (v *MOVLPS64LoadStore) Name() string {
	return "MOVLPS XMM (64 bit)"
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
	floats := [2]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [2]float32{}

	movlps64LoadStoreMovlps(&floats, &ret)

	log.Printf("MOVLPS64LoadStore input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVLPS64LoadStore) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
