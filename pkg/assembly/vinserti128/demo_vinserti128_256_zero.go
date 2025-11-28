package vinserti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vinserti128_256_zero.s
var assemblyVinserti128256Zero string

//go:embed stub_vinserti128_256_zero.go
var stubVinserti128256Zero string

type VINSERTI128256ZERO struct {
}

func (v *VINSERTI128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 16), // vals128 (4x u32)
		number.NewUintParameter(256, 32, 16), // vals256 (8x u32)
	}
}

func (v *VINSERTI128256ZERO) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VINSERTI128256ZERO) Name() string {
	return "VINSERTI128 (256 bit) zero"
}

func (v *VINSERTI128256ZERO) Description() string {
	return "Insert 128-bit block into lower 128-bit lane (0) of YMM; upper lane preserved from vals256."
}

func (v *VINSERTI128256ZERO) Stub() string {
	return stubVinserti128256Zero
}

func (v *VINSERTI128256ZERO) Assembly() string {
	return assemblyVinserti128256Zero
}

func (v *VINSERTI128256ZERO) Run(inputs [][]byte) (output []byte) {
	var vals128 [4]uint32
	copy(vals128[:], number.ToUint32Slice(inputs[0]))
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(inputs[1]))
	var ret [8]uint32
	vinserti128256Zero(&vals128, &vals256, &ret)
	log.Printf("VINSERTI128256 zero vals128 %v vals256 %v output %v", vals128, vals256, ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VINSERTI128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
