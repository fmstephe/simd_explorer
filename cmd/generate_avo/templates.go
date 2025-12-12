package main

const asmTemplate = `package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run {{.AssemblyGeneratorFileName}} -out ../{{.AssemblyFileName}} -stubs ../{{.StubFileName}} -pkg {{.PackageName}}
func main() {
	TEXT("{{.FunctionName}}", NOSPLIT, "func({{.Args}})")
	Comment("load params")
	{{.LoadArgsAvo}}

	{{.LoadRegistersAvo}}

	Comment("Sum packed float32 values from vals1 and vals2")
	{{.InstructionUpper}}(/* Add registers here */)


	{{.WriteReturnAvo}}

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
`
