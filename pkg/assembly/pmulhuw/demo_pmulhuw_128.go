package pmulhuw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmulhuw_128.s
var assemblyPmulhuw128 string

//go:embed stub_pmulhuw_128.go
var stubPmulhuw128 string

type PMULHUW128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewPMULHUW128() *PMULHUW128 {
	return &PMULHUW128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *PMULHUW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *PMULHUW128) Output() *number.Parameter {
	return v.ret
}

func (v *PMULHUW128) Name() string {
	return "PMULHUW (128 bit)"
}

func (v *PMULHUW128) Description() string {
	return "Packed multiply unsigned 16-bit integers; keep the high 16 bits of each 32-bit product."
}

func (v *PMULHUW128) Stub() string {
	return stubPmulhuw128
}

func (v *PMULHUW128) Assembly() string {
	return assemblyPmulhuw128
}

func (v *PMULHUW128) Run() (output []byte) {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	pmulhuw128(&vals1, &vals2, &ret)

	log.Printf("PMULHUW128 input %v %v output %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *PMULHUW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
