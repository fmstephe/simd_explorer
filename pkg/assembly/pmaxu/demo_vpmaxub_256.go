package pmaxu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxub_256.s
var assemblyVpmaxub256 string

//go:embed stub_vpmaxub_256.go
var stubVpmaxub256 string

type VPMAXUB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUB256() *VPMAXUB256 {
	return &VPMAXUB256{
		vals1: number.NewNamedUintParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPMAXUB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUB256) Name() string {
	return "VPMAXUB (256 bit) "
}

func (v *VPMAXUB256) Description() string {
	return "Unsigned maximum of packed 8-bit integers."
}

func (v *VPMAXUB256) Stub() string {
	return stubVpmaxub256
}

func (v *VPMAXUB256) Assembly() string {
	return assemblyVpmaxub256
}

func (v *VPMAXUB256) Run() {
	vals1 := [32]uint8{}
	copy(vals1[:], v.vals1.FlatData())
	vals2 := [32]uint8{}
	copy(vals2[:], v.vals2.FlatData())

	ret := [32]uint8{}

	vpmaxub256(&vals1, &vals2, &ret)

	log.Printf("VPMAXUB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	v.ret.SetData(ret[:])
}

func (v *VPMAXUB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
