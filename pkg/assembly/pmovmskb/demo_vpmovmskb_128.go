package pmovmskb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmovmskb_128.s
var assemblyVpmovmskb128 string

//go:embed stub_vpmovmskb_128.go
var stubVpmovmskb128 string

type VPMOVMSKB128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPMOVMSKB128() *VPMOVMSKB128 {
	return &VPMOVMSKB128{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VPMOVMSKB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPMOVMSKB128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMOVMSKB128) Name() string {
	return "VPMOVMSKB (128 bit)"
}

func (v *VPMOVMSKB128) Description() string {
	return "Move byte MSBs to mask (VEX): packs MSB of each of 16 bytes into a 16-bit mask (returned as 32-bit value)."
}

func (v *VPMOVMSKB128) Stub() string {
	return stubVpmovmskb128
}

func (v *VPMOVMSKB128) Assembly() string {
	return assemblyVpmovmskb128
}

func (v *VPMOVMSKB128) Run() (output []byte) {
	vals := [16]uint8{}
	copy(vals[:], v.vals.FlatData())

	var ret uint32

	vpmovmskb128(&vals, &ret)

	log.Printf("VPMOVMSKB128 input %v mask 0x%08x", vals, ret)

	out := number.Uint32ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VPMOVMSKB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
