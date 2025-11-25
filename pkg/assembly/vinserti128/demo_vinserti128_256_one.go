package vinserti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vinserti128_256_one.s
var assemblyVinserti128256One string

//go:embed stub_vinserti128_256_one.go
var stubVinserti128256One string

type VINSERTI128256ONE struct {
}

func (v *VINSERTI128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 16), // vals128 (4x u32)
		number.NewUintParameter(256, 32, 16), // vals256 (8x u32)
	}
}

func (v *VINSERTI128256ONE) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 16)
}

func (v *VINSERTI128256ONE) Name() string {
	return "VINSERTI128 (256 bit) one"
}

func (v *VINSERTI128256ONE) Description() string {
	return "Insert 128-bit block into upper 128-bit lane (1) of YMM; lower lane preserved from vals256."
}

func (v *VINSERTI128256ONE) Stub() string {
	return stubVinserti128256One
}

func (v *VINSERTI128256ONE) Assembly() string {
	return assemblyVinserti128256One
}

func (v *VINSERTI128256ONE) Run(inputs [][]byte) (output []byte) {
	var vals128 [4]uint32
	copy(vals128[:], number.ToUint32Slice(inputs[0]))
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(inputs[1]))
	var ret [8]uint32
	vinserti128256One(&vals128, &vals256, &ret)
	log.Printf("VINSERTI128256 one vals128 %v vals256 %v output %v", vals128, vals256, ret)
	return number.Uint32SliceToBytes(ret[:])
}

func (v *VINSERTI128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
