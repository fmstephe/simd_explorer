package vpbroadcastb

import (
	_ "embed"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpbroadcastb_512.s
var assemblyVpbroadcastb512 string

//go:embed stub_vpbroadcastb_512.go
var stubVpbroadcastb512 string

type VPBROADCASTB512 struct {
}

func (v *VPBROADCASTB512) Inputs() []*number.Parameter {
	return []*number.Parameter{number.NewUintParameter(8, 8, 16)}
}

func (v *VPBROADCASTB512) Output() *number.Parameter {
	return number.NewUintParameter(512, 64, 16)
}

func (v *VPBROADCASTB512) Name() string {
	return "VPBROADCASTB ZMM (512 bit)"
}

func (v *VPBROADCASTB512) Description() string {
	return "TODO"
}

func (v *VPBROADCASTB512) Stub() string {
	return stubVpbroadcastb512
}

func (v *VPBROADCASTB512) Assembly() string {
	return assemblyVpbroadcastb512
}

func (v *VPBROADCASTB512) Run(inputs [][]byte) (output []byte) {
	ret := [64]byte{}
	vpbroadcastb512(inputs[0][0], &ret)
	return ret[:]
}

func (v *VPBROADCASTB512) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
