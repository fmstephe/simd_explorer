package main

const asmTemplate = `package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run {{.AssemblyGeneratorFileName}} -out ../{{.AssemblyFileName}} -stubs ../{{.StubFileName}} -pkg {{.PackageName}}
func main() {
	// TODO: 
	// This code does NOT demonstrate the actual {{.InstructionUpper}} instruction.
	// Replace this implementation with the correct {{.InstructionUpper}} instruction code.
	TEXT("{{.FunctionName}}", NOSPLIT, "func(vals1, vals2, ret *[4]float32)")
	Comment("load params")
	vals1 := Load(Param("vals1"), GP64())
	vals2 := Load(Param("vals2"), GP64())
	ret := Load(Param("ret"), GP64())

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)
	
	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Sum packed float32 values from vals1 and vals2")
	ADDPS(regX2, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
`

const instructionDemoSource = `package {{.PackageName}}

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed {{.AssemblyFileName}}
var assembly{{.FunctionNameCamel}} string

//go:embed {{.StubFileName}}
var stub{{.FunctionNameCamel}} string

type {{.DemoTypeName}} struct {
	// TODO replace with actual parameters for instruction demo
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret *number.Parameter
}

func New{{.DemoTypeName}}() *{{.DemoTypeName}} {
	return &{{.DemoTypeName}} {
		// TODO replace with actual parameters for instruction demo
		vals1: number.NewNamedUintParameter("vals1", 128, 32, 10),
		vals2: number.NewNamedUintParameter("vals2", 128, 32, 10),
		ret: number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *{{.DemoTypeName}}) Inputs() []*number.Parameter {
	return []*number.Parameter {
		// TODO replace with actual parameters for instruction demo
		v.vals1,
		v.vals2,
	}
}

func (v *{{.DemoTypeName}}) Output() *number.Parameter {
	// TODO replace with actual parameters for instruction demo
	return v.ret
}

func (v *{{.DemoTypeName}}) Name() string {
	return "{{.InstructionUpper}} ({{.SizeClass}} bit) {{.Discriminator}}"
}

func (v *{{.DemoTypeName}}) Description() string {
	return "TODO add actual description of instruction being demoed"
}

func (v *{{.DemoTypeName}}) Stub() string {
	return stub{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Assembly() string {
	return assembly{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Run() {
	// TODO replace with actual parameters for instruction demo
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	{{.FunctionName}}(&vals1, &vals2, &ret)

	log.Printf("{{.DemoTypeName}} vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *{{.DemoTypeName}}) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
`
