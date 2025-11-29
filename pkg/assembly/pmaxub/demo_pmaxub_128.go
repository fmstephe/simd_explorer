package pmaxub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmaxub_128.s
var assemblyPmaxub128 string

//go:embed stub_pmaxub_128.go
var stubPmaxub128 string

type PMAXUB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPMAXUB128() *PMAXUB128 {
	return &PMAXUB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *PMAXUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PMAXUB128) Output() *number.Parameter {
	return v.ret
}

func (v *PMAXUB128) Name() string {
	return "PMAXUB (128 bit)"
}

func (v *PMAXUB128) Description() string {
	return "Packed max of unsigned bytes per lane."
}

func (v *PMAXUB128) Stub() string {
	return stubPmaxub128
}

func (v *PMAXUB128) Assembly() string {
	return assemblyPmaxub128
}

func (v *PMAXUB128) Run(_ [][]byte) (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	pmaxub128(&vals1, &vals2, &ret)

	log.Printf("PMAXUB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *PMAXUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
