package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubb_256.s
var assemblyVpsubb256 string

//go:embed stub_vpsubb_256.go
var stubVpsubb256 string

type VPSUBB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBB256() *VPSUBB256 {
	return &VPSUBB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSUBB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBB256) Name() string {
	return "VPSUBB (256 bit) "
}

func (v *VPSUBB256) Description() string {
	return "Subtract packed u8 bytes (wrap-around)."
}

func (v *VPSUBB256) Stub() string {
	return stubVpsubb256
}

func (v *VPSUBB256) Assembly() string {
	return assemblyVpsubb256
}

func (v *VPSUBB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpsubb256(&vals1, &vals2, &ret)

	log.Printf("VPSUBB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPSUBB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
