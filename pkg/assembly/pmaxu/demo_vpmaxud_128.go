package pmaxu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxud_128.s
var assemblyVpmaxud128 string

//go:embed stub_vpmaxud_128.go
var stubVpmaxud128 string

type VPMAXUD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUD128() *VPMAXUD128 {
	return &VPMAXUD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMAXUD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUD128) Name() string {
	return "VPMAXUD (128 bit) "
}

func (v *VPMAXUD128) Description() string {
	return "Unsigned maximum of packed 32-bit integers."
}

func (v *VPMAXUD128) Stub() string {
	return stubVpmaxud128
}

func (v *VPMAXUD128) Assembly() string {
	return assemblyVpmaxud128
}

func (v *VPMAXUD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpmaxud128(&vals1, &vals2, &ret)

	log.Printf("VPMAXUD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXUD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
