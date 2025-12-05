package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsb_128.s
var assemblyVpminsb128 string

//go:embed stub_vpminsb_128.go
var stubVpminsb128 string

type VPMINSB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSB128() *VPMINSB128 {
	return &VPMINSB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 8, 10),
	}
}

func (v *VPMINSB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSB128) Name() string {
	return "VPMINSB (128 bit) "
}

func (v *VPMINSB128) Description() string {
	return "Signed minimum of packed 8-bit integers."
}

func (v *VPMINSB128) Stub() string {
	return stubVpminsb128
}

func (v *VPMINSB128) Assembly() string {
	return assemblyVpminsb128
}

func (v *VPMINSB128) Run() {
	vals1 := [16]int8{}
	copy(vals1[:], number.BytesToInt8Slice(v.vals1.FlatData()))
	vals2 := [16]int8{}
	copy(vals2[:], number.BytesToInt8Slice(v.vals2.FlatData()))

	ret := [16]int8{}

	vpminsb128(&vals1, &vals2, &ret)

	log.Printf("VPMINSB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retSlc := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
}

func (v *VPMINSB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
