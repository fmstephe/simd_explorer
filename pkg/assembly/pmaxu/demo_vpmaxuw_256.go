package pmaxu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxuw_256.s
var assemblyVpmaxuw256 string

//go:embed stub_vpmaxuw_256.go
var stubVpmaxuw256 string

type VPMAXUW256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUW256() *VPMAXUW256 {
	return &VPMAXUW256{
		vals1: number.NewNamedUintParameter("vals1", 256, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPMAXUW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUW256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUW256) Name() string {
	return "VPMAXUW (256 bit) "
}

func (v *VPMAXUW256) Description() string {
	return "Unsigned maximum of packed 16-bit integers."
}

func (v *VPMAXUW256) Stub() string {
	return stubVpmaxuw256
}

func (v *VPMAXUW256) Assembly() string {
	return assemblyVpmaxuw256
}

func (v *VPMAXUW256) Run() {
	vals1 := [16]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [16]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [16]uint16{}

	vpmaxuw256(&vals1, &vals2, &ret)

	log.Printf("VPMAXUW256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXUW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
