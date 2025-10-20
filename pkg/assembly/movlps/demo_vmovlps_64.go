package movlps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovlps_64.s
var assemblyVmovlps64 string

//go:embed stub_vmovlps_64.go
var stubVmovlps64 string

type VMOVLPS64 struct {
}

func (v *VMOVLPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(64, 32),
		number.NewFloatParameter(64, 32),
	}
}

func (v *VMOVLPS64) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVLPS64) Name() string {
	return "VMOVLPS XMM (2X 64 bit)"
}

func (v *VMOVLPS64) Description() string {
	return "TODO"
}

func (v *VMOVLPS64) Stub() string {
	return stubVmovlps64
}

func (v *VMOVLPS64) Assembly() string {
	return assemblyVmovlps64
}

func (v *VMOVLPS64) Run(inputs [][]byte) (output []byte) {
	lower := [2]float32{}
	copy(lower[:], number.ToFloat32Slice(inputs[0]))

	upper := [2]float32{}
	copy(upper[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vmovlps64(&lower, &upper, &ret)

	log.Printf("VMOVLPS64 input lower %v upper %v output %v", lower, upper, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVLPS64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
