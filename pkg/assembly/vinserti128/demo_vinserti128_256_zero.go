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
	vals128 *number.Parameter
	vals256 *number.Parameter
	ret     *number.Parameter
}

func NewVINSERTI128256ZERO() *VINSERTI128256ZERO {
	return &VINSERTI128256ZERO{
		vals128: number.NewNamedUintParameter("vals128", 128, 32, 10),
		vals256: number.NewNamedUintParameter("vals256", 256, 32, 10),
		ret:     number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VINSERTI128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals128,
		v.vals256,
	}
}

func (v *VINSERTI128256ZERO) Output() *number.Parameter {
	return v.ret
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

func (v *VINSERTI128256ZERO) Run() {
	var vals128 [4]uint32
	copy(vals128[:], number.ToUint32Slice(v.vals128.FlatData()))
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(v.vals256.FlatData()))
	var ret [8]uint32
	vinserti128256Zero(&vals128, &vals256, &ret)
	log.Printf("VINSERTI128256 zero vals128 %v vals256 %v output %v", vals128, vals256, ret)
	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VINSERTI128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
