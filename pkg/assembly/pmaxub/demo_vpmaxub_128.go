package pmaxub

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxub_128.s
var assemblyVpmaxub128 string

//go:embed stub_vpmaxub_128.go
var stubVpmaxub128 string

type VPMAXUB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUB128() *VPMAXUB128 {
	return &VPMAXUB128{
		vals1: number.NewNamedUintParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPMAXUB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUB128) Name() string {
	return "VPMAXUB (128 bit)"
}

func (v *VPMAXUB128) Description() string {
	return "Packed max of unsigned bytes per lane (VEX)."
}

func (v *VPMAXUB128) Stub() string {
	return stubVpmaxub128
}

func (v *VPMAXUB128) Assembly() string {
	return assemblyVpmaxub128
}

func (v *VPMAXUB128) Run() (output []byte) {
	vals1 := [16]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [16]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [16]uint8{}

	vpmaxub128(&vals1, &vals2, &ret)

	log.Printf("VPMAXUB128 input %v %v output %v", vals1, vals2, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *VPMAXUB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
