package pmovmskb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_pmovmskb_128.s
var assemblyPmovmskb128 string

//go:embed stub_pmovmskb_128.go
var stubPmovmskb128 string

type PMOVMSKB128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewPMOVMSKB128() *PMOVMSKB128 {
	return &PMOVMSKB128{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *PMOVMSKB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *PMOVMSKB128) Output() *number.Parameter {
	return v.ret
}

func (v *PMOVMSKB128) Name() string {
	return "PMOVMSKB (128 bit)"
}

func (v *PMOVMSKB128) Description() string {
	return "Move byte MSBs to mask: packs the most-significant bit of each of 16 bytes into a 16-bit mask (returned as 32-bit value)."
}

func (v *PMOVMSKB128) Stub() string {
	return stubPmovmskb128
}

func (v *PMOVMSKB128) Assembly() string {
	return assemblyPmovmskb128
}

func (v *PMOVMSKB128) Run() {
	vals := [16]uint8{}
	copy(vals[:], v.vals.FlatData())

	var ret uint32

	pmovmskb128(&vals, &ret)

	log.Printf("PMOVMSKB128 input %v mask 0x%08x", vals, ret)

	out := number.Uint32ToBytes(ret)
	v.ret.SetData(out)

}

func (v *PMOVMSKB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
