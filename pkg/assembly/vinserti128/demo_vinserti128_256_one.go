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
	vals128 *number.Parameter
	vals256 *number.Parameter
	ret     *number.Parameter
}

func NewVINSERTI128256ONE() *VINSERTI128256ONE {
	return &VINSERTI128256ONE{
		vals128: number.NewNamedUintParameter("vals128", 128, 32, 10),
		vals256: number.NewNamedUintParameter("vals256", 256, 32, 10),
		ret:     number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VINSERTI128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals128,
		v.vals256,
	}
}

func (v *VINSERTI128256ONE) Output() *number.Parameter {
	return v.ret
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

func (v *VINSERTI128256ONE) Run() {
	vals128 := [4]uint32{}
	copy(vals128[:], number.ToUint32Slice(v.vals128.FlatData()))
	vals256 := [8]uint32{}
	copy(vals256[:], number.ToUint32Slice(v.vals256.FlatData()))
	ret := [8]uint32{}
	copy(ret[:], number.ToUint32Slice(v.ret.FlatData()))

	vinserti128256One(&vals128, &vals256, &ret)

	log.Printf("VINSERTI128256 vals128 %v vals256 %v ret %v", vals128, vals256, ret)

	retBytes := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VINSERTI128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
