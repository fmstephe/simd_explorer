package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_256_k.s
var assembly256K string

//go:embed stub_256_k.go
var stub256K string

type Vpbroadcastb256K struct {
}

func (v *Vpbroadcastb256K) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(8, 8, 16),
		number.NewUintParameter(64, 64, 16),
	}
}

func (v *Vpbroadcastb256K) Output() *number.Parameter {
	return number.NewUintParameter(256, 64, 16)
}

func (v *Vpbroadcastb256K) Name() string {
	return "VPBROADCASTB YMM (256K bit) K"
}

func (v *Vpbroadcastb256K) Description() string {
	return "TODO"
}

func (v *Vpbroadcastb256K) Stub() string {
	return stub256K
}

func (v *Vpbroadcastb256K) Assembly() string {
	return assembly256K
}

func (v *Vpbroadcastb256K) Run(inputs [][]byte) (output []byte) {
	ret := [32]byte{}
	vpbroadcastb256K(inputs[0][0], number.ToUint64(inputs[1]), &ret)
	return ret[:]
}

func (v *Vpbroadcastb256K) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
