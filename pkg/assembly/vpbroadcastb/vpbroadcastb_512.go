package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
)

//go:embed asm_512.s
var assembly512 string

//go:embed stub_512.go
var stub512 string

type Vpbroadcastb512 struct {
}

func (v *Vpbroadcastb512) InputSizes() []int {
	return []int{8}
}

func (v *Vpbroadcastb512) OutputSize() int {
	return 512
}

func (v *Vpbroadcastb512) Name() string {
	return "VPBROADCASTB ZMM (512 bit)"
}

func (v *Vpbroadcastb512) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb512) Stub() string {
	return stub512
}

func (v *Vpbroadcastb512) Assembly() string {
	return assembly512
}

func (v *Vpbroadcastb512) Run(inputs [][]byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512(inputs[0][0], &ret)
	return ret[:]
}

func (v *Vpbroadcastb512) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
