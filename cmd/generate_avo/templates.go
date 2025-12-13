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
{{.LoadArgsAvo}}

{{.LoadRegistersAvo}}

	Comment("Execute the instruction being demonstrated")
	{{.InstructionUpper}}(/* Add registers here */)


{{.WriteReturnAvo}}

{{.VZeroUpperAvo}}

	Comment("Return from function")
	RET()

	// generate!
	Generate()
}
`
