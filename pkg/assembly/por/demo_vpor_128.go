package por

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpor_128.s
var assemblyVpor128 string

//go:embed stub_vpor_128.go
var stubVpor128 string

type VPOR128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPOR128() *VPOR128 {
	return &VPOR128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 16),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 16),
	}
}

func (v *VPOR128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPOR128) Output() *number.Parameter {
	return v.ret
}

func (v *VPOR128) Name() string {
	return "VPOR (128 bit) "
}

func (v *VPOR128) Description() string {
	return "Bitwise OR of packed bytes."
}

func (v *VPOR128) Stub() string {
	return stubVpor128
}

func (v *VPOR128) Assembly() string {
	return assemblyVpor128
}

func (v *VPOR128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpor128(&vals1, &vals2, &ret)

	log.Printf("VPOR128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPOR128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
