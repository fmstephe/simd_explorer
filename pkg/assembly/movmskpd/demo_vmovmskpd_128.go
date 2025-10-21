package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskpd_128.s
var assemblyVmovmskpd128 string

//go:embed stub_vmovmskpd_128.go
var stubVmovmskpd128 string

type VMOVMSKPD128 struct {
}

func (v *VMOVMSKPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 64),
	}
}

func (v *VMOVMSKPD128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *VMOVMSKPD128) Name() string {
	return "VMOVMSKPD (128 bit) "
}

func (v *VMOVMSKPD128) Description() string {
	return "TODO"
}

func (v *VMOVMSKPD128) Stub() string {
	return stubVmovmskpd128
}

func (v *VMOVMSKPD128) Assembly() string {
	return assemblyVmovmskpd128
}

func (v *VMOVMSKPD128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [2]float64{}
	copy(floats[:], number.ToFloat64Slice(inputs[0]))

	ret := [4]byte{}

	vmovmskpd128(&floats, &ret)

	log.Printf("VMOVMSKPD128 input %v output %v", floats, ret)

	return ret[:]
}

func (v *VMOVMSKPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
