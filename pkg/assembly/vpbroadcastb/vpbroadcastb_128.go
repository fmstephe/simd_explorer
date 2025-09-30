package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128.s
var assembly128 string

//go:embed stub_128.go
var stub128 string

type Vpbroadcastb128 struct {
}

func (v *Vpbroadcastb128) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(8, 8, 16)}
}

func (v *Vpbroadcastb128) Output() *number.Parameter {
	return number.NewUintParameter(128, 64, 16)
}

func (v *Vpbroadcastb128) Name() string {
	return "VPBROADCASTB XMM (128 bit)"
}

func (v *Vpbroadcastb128) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb128) Stub() string {
	return stub128
}

func (v *Vpbroadcastb128) Assembly() string {
	return assembly128
}

func (v *Vpbroadcastb128) Run(inputs [][]byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128(inputs[0][0], &ret)
	return ret[:]
}

func (v *Vpbroadcastb128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
