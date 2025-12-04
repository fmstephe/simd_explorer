package psub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsubb_128.s
var assemblyVpsubb128 string

//go:embed stub_vpsubb_128.go
var stubVpsubb128 string

type VPSUBB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPSUBB128() *VPSUBB128 {
	return &VPSUBB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSUBB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPSUBB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPSUBB128) Name() string {
	return "VPSUBB (128 bit) "
}

func (v *VPSUBB128) Description() string {
	return "Subtract packed u8 bytes (wrap-around)."
}

func (v *VPSUBB128) Stub() string {
	return stubVpsubb128
}

func (v *VPSUBB128) Assembly() string {
	return assemblyVpsubb128
}

func (v *VPSUBB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpsubb128(&vals1, &vals2, &ret)

	log.Printf("VPSUBB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPSUBB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
