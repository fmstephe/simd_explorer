package pavgb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pavgb_128.s
var assemblyPavgb128 string

//go:embed stub_pavgb_128.go
var stubPavgb128 string

type PAVGB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPAVGB128() *PAVGB128 {
	return &PAVGB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *PAVGB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PAVGB128) Output() *number.Parameter {
	return v.ret
}

func (v *PAVGB128) Name() string {
	return "PAVGB (128 bit)"
}

func (v *PAVGB128) Description() string {
	return "Average of packed unsigned bytes with rounding: (a+b+1)>>1."
}

func (v *PAVGB128) Stub() string {
	return stubPavgb128
}

func (v *PAVGB128) Assembly() string {
	return assemblyPavgb128
}

func (v *PAVGB128) Run() (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	pavgb128(&vals1, &vals2, &ret)

	log.Printf("PAVGB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *PAVGB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
