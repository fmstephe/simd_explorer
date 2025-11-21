package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_128_reverse.s
var assemblyVpermilps128Reverse string

//go:embed stub_vpermilps_128_reverse.go
var stubVpermilps128Reverse string

type VPERMILPS128REVERSE struct {
}

func (v *VPERMILPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VPERMILPS128REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VPERMILPS128REVERSE) Name() string {
	return "VPERMILPS (128 bit) reverse"
}

func (v *VPERMILPS128REVERSE) Description() string {
	return "Permute with imm8=0x1B: reverse lane order [a3 a2 a1 a0]."
}

func (v *VPERMILPS128REVERSE) Stub() string {
	return stubVpermilps128Reverse
}

func (v *VPERMILPS128REVERSE) Assembly() string {
	return assemblyVpermilps128Reverse
}

func (v *VPERMILPS128REVERSE) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vpermilps128Reverse(&vals, &ret)

	log.Printf("VPERMILPS128REVERSE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
