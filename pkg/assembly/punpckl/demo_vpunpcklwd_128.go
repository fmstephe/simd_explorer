package punpckl

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpunpcklwd_128.s
var assemblyVpunpcklwd128 string

//go:embed stub_vpunpcklwd_128.go
var stubVpunpcklwd128 string

type VPUNPCKLWD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPUNPCKLWD128() *VPUNPCKLWD128 {
	return &VPUNPCKLWD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 16, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 16, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPUNPCKLWD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPUNPCKLWD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPUNPCKLWD128) Name() string {
	return "VPUNPCKLWD (128 bit) "
}

func (v *VPUNPCKLWD128) Description() string {
	return "Unpack and interleave low-order words from two 128-bit sources."
}

func (v *VPUNPCKLWD128) Stub() string {
	return stubVpunpcklwd128
}

func (v *VPUNPCKLWD128) Assembly() string {
	return assemblyVpunpcklwd128
}

func (v *VPUNPCKLWD128) Run() {
	vals1 := [8]uint16{}
	copy(vals1[:], number.ToUint16Slice(v.vals1.FlatData()))
	vals2 := [8]uint16{}
	copy(vals2[:], number.ToUint16Slice(v.vals2.FlatData()))

	ret := [8]uint16{}

	vpunpcklwd128(&vals1, &vals2, &ret)

	log.Printf("VPUNPCKLWD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPUNPCKLWD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
