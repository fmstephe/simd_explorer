package pavgw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pabgw_128.s
var assemblyPabgw128 string

//go:embed stub_pabgw_128.go
var stubPabgw128 string

type PAVGW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPAVGW128() *PAVGW128 {
	return &PAVGW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *PAVGW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PAVGW128) Output() *number.Parameter {
	return v.ret
}

func (v *PAVGW128) Name() string {
	return "PAVGW (128 bit)"
}

func (v *PAVGW128) Description() string {
	return "Average of packed unsigned 16-bit words with rounding: (a+b+1)>>1."
}

func (v *PAVGW128) Stub() string {
	return stubPabgw128
}

func (v *PAVGW128) Assembly() string {
	return assemblyPabgw128
}

func (v *PAVGW128) Run() (output []byte) {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	pabgw128(&vals1, &vals2, &ret)

	log.Printf("PAVGW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *PAVGW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
