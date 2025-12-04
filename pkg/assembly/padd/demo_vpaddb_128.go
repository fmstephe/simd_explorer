package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddb_128.s
var assemblyVpaddb128 string

//go:embed stub_vpaddb_128.go
var stubVpaddb128 string

type VPADDB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDB128() *VPADDB128 {
	return &VPADDB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPADDB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDB128) Name() string {
	return "VPADDB (128 bit) "
}

func (v *VPADDB128) Description() string {
	return "Add packed u8 bytes (wrap-around)."
}

func (v *VPADDB128) Stub() string {
	return stubVpaddb128
}

func (v *VPADDB128) Assembly() string {
	return assemblyVpaddb128
}

func (v *VPADDB128) Run() {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpaddb128(&vals1, &vals2, &ret)

	log.Printf("VPADDB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
}

func (v *VPADDB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
