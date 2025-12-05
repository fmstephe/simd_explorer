package pmins

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpminsb_256.s
var assemblyVpminsb256 string

//go:embed stub_vpminsb_256.go
var stubVpminsb256 string

type VPMINSB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMINSB256() *VPMINSB256 {
	return &VPMINSB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 8, 10),
	}
}

func (v *VPMINSB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMINSB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMINSB256) Name() string {
	return "VPMINSB (256 bit) "
}

func (v *VPMINSB256) Description() string {
	return "Signed minimum of packed 8-bit integers."
}

func (v *VPMINSB256) Stub() string {
	return stubVpminsb256
}

func (v *VPMINSB256) Assembly() string {
	return assemblyVpminsb256
}

func (v *VPMINSB256) Run() {
	vals1 := [32]int8{}
	copy(vals1[:], number.BytesToInt8Slice(v.vals1.FlatData()))
	vals2 := [32]int8{}
	copy(vals2[:], number.BytesToInt8Slice(v.vals2.FlatData()))

	ret := [32]int8{}

	vpminsb256(&vals1, &vals2, &ret)

	log.Printf("VPMINSB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retSlc := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
}

func (v *VPMINSB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
