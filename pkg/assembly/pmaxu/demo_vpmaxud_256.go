package pmaxu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxud_256.s
var assemblyVpmaxud256 string

//go:embed stub_vpmaxud_256.go
var stubVpmaxud256 string

type VPMAXUD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUD256() *VPMAXUD256 {
	return &VPMAXUD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMAXUD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUD256) Name() string {
	return "VPMAXUD (256 bit) "
}

func (v *VPMAXUD256) Description() string {
	return "Unsigned maximum of packed 32-bit integers."
}

func (v *VPMAXUD256) Stub() string {
	return stubVpmaxud256
}

func (v *VPMAXUD256) Assembly() string {
	return assemblyVpmaxud256
}

func (v *VPMAXUD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpmaxud256(&vals1, &vals2, &ret)

	log.Printf("VPMAXUD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXUD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
