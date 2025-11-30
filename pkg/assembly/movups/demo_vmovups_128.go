package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovups_128.s
var assemblyVmovups128 string

//go:embed stub_vmovups_128.go
var stubVmovups128 string

type VMOVUPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVUPS128() *VMOVUPS128 {
	return &VMOVUPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMOVUPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVUPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVUPS128) Name() string {
	return "VMOVUPS XMM (128 bit)"
}

func (v *VMOVUPS128) Description() string {
	return "Unaligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *VMOVUPS128) Stub() string {
	return stubVmovups128
}

func (v *VMOVUPS128) Assembly() string {
	return assemblyVmovups128
}

func (v *VMOVUPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vmovups128(&vals, &ret)

	log.Printf("VMOVUPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMOVUPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
