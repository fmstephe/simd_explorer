package main

const asmTemplate = `package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run {{.AssemblyGeneratorFileName}} -out ../{{.AssemblyFileName}} -stubs ../{{.StubFileName}} -pkg {{.PackageName}}
func main() {
	// NOTE: This is a generic template that generates valid assembly for testing purposes.
	// It does NOT implement the actual {{.InstructionUpper}} instruction.
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
}

func (v *{{.DemoTypeName}}) Inputs() []*number.Parameter {
	// TODO
	return nil
}

func (v *{{.DemoTypeName}}) Output() *number.Parameter {
	// TODO
	return nil
}

func (v *{{.DemoTypeName}}) Name() string {
	return "{{.InstructionUpper}} ({{.SizeClass}} bit) {{.Discriminator}}"
}

func (v *{{.DemoTypeName}}) Description() string {
	return "TODO"
}

func (v *{{.DemoTypeName}}) Stub() string {
	return stub{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Assembly() string {
	return assembly{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	{{.FunctionName}}(/* TODO */)

	log.Printf("{{.DemoTypeName}} input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *{{.DemoTypeName}}) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
`
