package pminub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pminub_128.s
var assemblyPminub128 string

//go:embed stub_pminub_128.go
var stubPminub128 string

type PMINUB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPMINUB128() *PMINUB128 {
	return &PMINUB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *PMINUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PMINUB128) Output() *number.Parameter {
	return v.ret
}

func (v *PMINUB128) Name() string {
	return "PMINUB (128 bit)"
}

func (v *PMINUB128) Description() string {
	return "Packed min of unsigned bytes per lane."
}

func (v *PMINUB128) Stub() string {
	return stubPminub128
}

func (v *PMINUB128) Assembly() string {
	return assemblyPminub128
}

func (v *PMINUB128) Run() (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	pminub128(&vals1, &vals2, &ret)

	log.Printf("PMINUB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *PMINUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
