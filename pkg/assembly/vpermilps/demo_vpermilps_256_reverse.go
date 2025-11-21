package vpermilps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilps_256_reverse.s
var assemblyVpermilps256Reverse string

//go:embed stub_vpermilps_256_reverse.go
var stubVpermilps256Reverse string

type VPERMILPS256REVERSE struct {
}

func (v *VPERMILPS256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VPERMILPS256REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VPERMILPS256REVERSE) Name() string {
	return "VPERMILPS (256 bit) reverse"
}

func (v *VPERMILPS256REVERSE) Description() string {
	return "Permute with imm8=0x1B per 128-bit lane: reverse lane order."
}

func (v *VPERMILPS256REVERSE) Stub() string {
	return stubVpermilps256Reverse
}

func (v *VPERMILPS256REVERSE) Assembly() string {
	return assemblyVpermilps256Reverse
}

func (v *VPERMILPS256REVERSE) Run(inputs [][]byte) (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vpermilps256Reverse(&vals, &ret)

	log.Printf("VPERMILPS256REVERSE vals %v ret %v", vals, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VPERMILPS256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
