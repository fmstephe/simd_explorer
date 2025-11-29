package pmaxsw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaxsw_128.s
var assemblyVpmaxsw128 string

//go:embed stub_vpmaxsw_128.go
var stubVpmaxsw128 string

type VPMAXSW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPMAXSW128() *VPMAXSW128 {
	return &VPMAXSW128{
		vals1: number.NewNamedIntParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedIntParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedIntParameter("ret", 128, 16, 10),
	}
}

func (v *VPMAXSW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPMAXSW128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMAXSW128) Name() string {
	return "VPMAXSW (128 bit)"
}

func (v *VPMAXSW128) Description() string {
	return "Packed max of signed 16-bit words per lane (VEX)."
}

func (v *VPMAXSW128) Stub() string {
	return stubVpmaxsw128
}

func (v *VPMAXSW128) Assembly() string {
	return assemblyVpmaxsw128
}

func (v *VPMAXSW128) Run(_ [][]byte) (output []byte) {
	vals1 := [8]int16{}
	copy(vals1[:], number.ToInt16Slice(v.vals1.FlatData()))
	vals2 := [8]int16{}
	copy(vals2[:], number.ToInt16Slice(v.vals2.FlatData()))

	ret := [8]int16{}

	vpmaxsw128(&vals1, &vals2, &ret)

	log.Printf("VPMAXSW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Int16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPMAXSW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
