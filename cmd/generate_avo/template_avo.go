package main

const avoTemplate = `package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

//go:generate go run {{.AssemblyGeneratorFileName}} -out ../{{.AssemblyFileName}} -stubs ../{{.StubFileName}} -pkg {{.PackageName}}
func main() {
	TEXT("{{.FunctionName}}", NOSPLIT, "func({{.Args}})")
	Comment("load params")
{{.AvoLoadArgs}}

{{.AvoLoadRegisters}}

	Comment("Execute the instruction being demonstrated")
	{{.InstructionUpper}}(/* Add registers here */)


{{.AvoWriteReturn}}

{{.AvoVZeroUpper}}

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
`
