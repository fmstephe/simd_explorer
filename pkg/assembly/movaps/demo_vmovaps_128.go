package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovaps_128.s
var assemblyVmovaps128 string

//go:embed stub_vmovaps_128.go
var stubVmovaps128 string

type VMOVAPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVAPS128() *VMOVAPS128 {
	return &VMOVAPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMOVAPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVAPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVAPS128) Name() string {
	return "VMOVAPS XMM (128 bit)"
}

func (v *VMOVAPS128) Description() string {
	return "Aligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *VMOVAPS128) Stub() string {
	return stubVmovaps128
}

func (v *VMOVAPS128) Assembly() string {
	return assemblyVmovaps128
}

func (v *VMOVAPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vmovaps128(&vals, &ret)

	log.Printf("VMOVAPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMOVAPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
