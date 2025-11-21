package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_256_identity.s
var assemblyVpermilpd256Identity string

//go:embed stub_vpermilpd_256_identity.go
var stubVpermilpd256Identity string

type VPERMILPD256IDENTITY struct {
}

func (v *VPERMILPD256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 64),
	}
}

func (v *VPERMILPD256IDENTITY) Output() *number.Parameter {
	return number.NewFloatParameter(256, 64)
}

func (v *VPERMILPD256IDENTITY) Name() string {
	return "VPERMILPD (256 bit) identity"
}

func (v *VPERMILPD256IDENTITY) Description() string {
	return "Permute double-precision floats with imm8=0xE4 per 128-bit lane: identity."
}

func (v *VPERMILPD256IDENTITY) Stub() string {
	return stubVpermilpd256Identity
}

func (v *VPERMILPD256IDENTITY) Assembly() string {
	return assemblyVpermilpd256Identity
}

func (v *VPERMILPD256IDENTITY) Run(inputs [][]byte) (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(inputs[0]))

	ret := [4]float64{}

	vpermilpd256Identity(&vals, &ret)

	log.Printf("VPERMILPD256IDENTITY vals %v ret %v", vals, ret)

	return number.Float64SliceToBytes(ret[:])
}

func (v *VPERMILPD256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
