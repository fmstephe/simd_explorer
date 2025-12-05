package pmaxs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsb_256.s
var assemblyVpmaxsb256 string

//go:embed stub_vpmaxsb_256.go
var stubVpmaxsb256 string

type VPMAXSB256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSB256() *VPMAXSB256 {
	return &VPMAXSB256{
		vals1: number.NewNamedIntParameter("vals1", 256, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 256, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 256, 8, 10),
	}
}

func (v *VPMAXSB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSB256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSB256) Name() string {
	return "VPMAXSB (256 bit) "
}

func (v *VPMAXSB256) Description() string {
	return "Signed maximum of packed 8-bit integers."
}

func (v *VPMAXSB256) Stub() string {
	return stubVpmaxsb256
}

func (v *VPMAXSB256) Assembly() string {
	return assemblyVpmaxsb256
}

func (v *VPMAXSB256) Run() {
	vals1 := [32]int8{}
	copy(vals1[:], number.BytesToInt8Slice(v.vals1.FlatData()))
	vals2 := [32]int8{}
	copy(vals2[:], number.BytesToInt8Slice(v.vals2.FlatData()))

	ret := [32]int8{}

	vpmaxsb256(&vals1, &vals2, &ret)

	log.Printf("VPMAXSB256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retSlc := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
}

func (v *VPMAXSB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
