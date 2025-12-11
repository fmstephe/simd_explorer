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
	TEXT("{{.FunctionName}}", NOSPLIT, "func({{.Args}})")
	Comment("load params")
	{{.LoadParams}}

	Comment("Load vals1 into XMM register")
	regX1 := XMM()
	VMOVDQA(Mem{Base: vals1}, regX1)
	
	Comment("Load vals2 into XMM register")
	regX2 := XMM()
	VMOVDQA(Mem{Base: vals2}, regX2)

	Comment("Sum packed float32 values from vals1 and vals2")
	{{.InstructionUpper}}(regX2, regX1)

	Comment("Write results into return memory address")
	VMOVDQA(regX1, Mem{Base: ret})

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
`
