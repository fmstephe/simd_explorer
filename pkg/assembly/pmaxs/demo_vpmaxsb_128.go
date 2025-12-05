package pmaxs

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsb_128.s
var assemblyVpmaxsb128 string

//go:embed stub_vpmaxsb_128.go
var stubVpmaxsb128 string

type VPMAXSB128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSB128() *VPMAXSB128 {
	return &VPMAXSB128{
		vals1: number.NewNamedIntParameter("vals1", 128, 8, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 8, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 8, 10),
	}
}

func (v *VPMAXSB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSB128) Name() string {
	return "VPMAXSB (128 bit) "
}

func (v *VPMAXSB128) Description() string {
	return "Signed maximum of packed 8-bit integers."
}

func (v *VPMAXSB128) Stub() string {
	return stubVpmaxsb128
}

func (v *VPMAXSB128) Assembly() string {
	return assemblyVpmaxsb128
}

func (v *VPMAXSB128) Run() {
	vals1 := [16]int8{}
	copy(vals1[:], number.BytesToInt8Slice(v.vals1.FlatData()))
	vals2 := [16]int8{}
	copy(vals2[:], number.BytesToInt8Slice(v.vals2.FlatData()))

	ret := [16]int8{}

	vpmaxsb128(&vals1, &vals2, &ret)

	log.Printf("VPMAXSB128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retSlc := number.Int8SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
}

func (v *VPMAXSB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
