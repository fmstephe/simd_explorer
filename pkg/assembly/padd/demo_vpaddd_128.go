package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddd_128.s
var assemblyVpaddd128 string

//go:embed stub_vpaddd_128.go
var stubVpaddd128 string

type VPADDD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDD128() *VPADDD128 {
	return &VPADDD128{
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPADDD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDD128) Name() string {
	return "VPADDD (128 bit) "
}

func (v *VPADDD128) Description() string {
	return "Add packed u32 dwords (wrap-around)."
}

func (v *VPADDD128) Stub() string {
	return stubVpaddd128
}

func (v *VPADDD128) Assembly() string {
	return assemblyVpaddd128
}

func (v *VPADDD128) Run() {
	vals1 := [4]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [4]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [4]uint32{}

	vpaddd128(&vals1, &vals2, &ret)

	log.Printf("VPADDD128 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
