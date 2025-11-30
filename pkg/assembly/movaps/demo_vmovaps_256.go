package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovaps_256.s
var assemblyVmovaps256 string

//go:embed stub_vmovaps_256.go
var stubVmovaps256 string

type VMOVAPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVAPS256() *VMOVAPS256 {
	return &VMOVAPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMOVAPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVAPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVAPS256) Name() string {
	return "VMOVAPS YMM (256 bit)"
}

func (v *VMOVAPS256) Description() string {
	return "Aligned move of packed single-precision floats between memory and YMM; copies data unchanged."
}

func (v *VMOVAPS256) Stub() string {
	return stubVmovaps256
}

func (v *VMOVAPS256) Assembly() string {
	return assemblyVmovaps256
}

func (v *VMOVAPS256) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vmovaps256(&vals, &ret)

	log.Printf("VMOVAPS256 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVAPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
