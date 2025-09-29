package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_k.s
var assembly128K string

//go:embed stub_128_k.go
var stub128K string

type Vpbroadcastb128K struct {
}

func (v *Vpbroadcastb128K) InputSizes() []number.Converter {
	return []number.Converter{number.NewUintConverter(8, 16)}
}

func (v *Vpbroadcastb128K) OutputSize() number.Converter {
	return number.NewUintConverter(128, 16)
}

func (v *Vpbroadcastb128K) Name() string {
	return "VPBROADCASTB XMM (128K bit) K"
}

func (v *Vpbroadcastb128K) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb128K) Stub() string {
	return stub128K
}

func (v *Vpbroadcastb128K) Assembly() string {
	return assembly128K
}

func (v *Vpbroadcastb128K) Run(inputs [][]byte) (output []byte) {
	ret := [16]byte{}
	vpbroadcastb128K(inputs[0][0], number.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *Vpbroadcastb128K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
