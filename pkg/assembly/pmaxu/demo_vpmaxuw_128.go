package pmaxu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxuw_128.s
var assemblyVpmaxuw128 string

//go:embed stub_vpmaxuw_128.go
var stubVpmaxuw128 string

type VPMAXUW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXUW128() *VPMAXUW128 {
	return &VPMAXUW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPMAXUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXUW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXUW128) Name() string {
	return "VPMAXUW (128 bit) "
}

func (v *VPMAXUW128) Description() string {
	return "Unsigned maximum of packed 16-bit integers."
}

func (v *VPMAXUW128) Stub() string {
	return stubVpmaxuw128
}

func (v *VPMAXUW128) Assembly() string {
	return assemblyVpmaxuw128
}

func (v *VPMAXUW128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpmaxuw128(&vals1, &vals2, &ret)

	log.Printf("VPMAXUW128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMAXUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
