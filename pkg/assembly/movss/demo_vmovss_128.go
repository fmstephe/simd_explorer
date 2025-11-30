package movss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovss_128.s
var assemblyVmovss128 string

//go:embed stub_vmovss_128.go
var stubVmovss128 string

type VMOVSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVSS128() *VMOVSS128 {
	return &VMOVSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMOVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVSS128) Name() string {
	return "VMOVSS XMM (128 bit)"
}

func (v *VMOVSS128) Description() string {
	return "AVX form: move scalar single-precision (lane 0) between XMM and memory; upper lanes pass through."
}

func (v *VMOVSS128) Stub() string {
	return stubVmovss128
}

func (v *VMOVSS128) Assembly() string {
	return assemblyVmovss128
}

func (v *VMOVSS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vmovss128(&vals, &ret)

	log.Printf("VMOVSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
