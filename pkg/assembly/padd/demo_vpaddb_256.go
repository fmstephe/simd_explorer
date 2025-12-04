package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddb_256.s
var assemblyVpaddb256 string

//go:embed stub_vpaddb_256.go
var stubVpaddb256 string

type VPADDB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDB256() *VPADDB256 {
	return &VPADDB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPADDB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDB256) Name() string {
	return "VPADDB (256 bit) "
}

func (v *VPADDB256) Description() string {
	return "Add packed u8 bytes (wrap-around)."
}

func (v *VPADDB256) Stub() string {
	return stubVpaddb256
}

func (v *VPADDB256) Assembly() string {
	return assemblyVpaddb256
}

func (v *VPADDB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpaddb256(&vals1, &vals2, &ret)

	log.Printf("VPADDB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
}

func (v *VPADDB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
