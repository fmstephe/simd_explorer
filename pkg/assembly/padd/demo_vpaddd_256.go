package padd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpaddd_256.s
var assemblyVpaddd256 string

//go:embed stub_vpaddd_256.go
var stubVpaddd256 string

type VPADDD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVPADDD256() *VPADDD256 {
	return &VPADDD256{
		vals1: number.NewNamedUintParameter("vals1", 256, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 256, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPADDD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VPADDD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPADDD256) Name() string {
	return "VPADDD (256 bit) "
}

func (v *VPADDD256) Description() string {
	return "Add packed u32 dwords (wrap-around)."
}

func (v *VPADDD256) Stub() string {
	return stubVpaddd256
}

func (v *VPADDD256) Assembly() string {
	return assemblyVpaddd256
}

func (v *VPADDD256) Run() {
	vals1 := [8]uint32{}
	copy(vals1[:], number.ToUint32Slice(v.vals1.FlatData()))
	vals2 := [8]uint32{}
	copy(vals2[:], number.ToUint32Slice(v.vals2.FlatData()))

	ret := [8]uint32{}

	vpaddd256(&vals1, &vals2, &ret)

	log.Printf("VPADDD256 vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPADDD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
