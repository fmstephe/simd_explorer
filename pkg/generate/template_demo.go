package generate

const demoTemplate = `package {{.PackageName}}

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
{{.DemoFields}}
}

func New{{.DemoTypeName}}() *{{.DemoTypeName}} {
	return &{{.DemoTypeName}} {
{{.DemoConstructor}}
	}
}

func (v *{{.DemoTypeName}}) Inputs() []*number.Parameter {
	return []*number.Parameter {
{{.DemoInputs}}	}
}

func (v *{{.DemoTypeName}}) Output() *number.Parameter {
	return v.ret
}

func (v *{{.DemoTypeName}}) Name() string {
	return "{{.DemoName}}"
}

func (v *{{.DemoTypeName}}) Description() string {
	return {{.DemoDescription}}
}

func (v *{{.DemoTypeName}}) Stub() string {
	return stub{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Assembly() string {
	return assembly{{.FunctionNameCamel}}
}

func (v *{{.DemoTypeName}}) Run() {
{{.DemoInitArrays}}
	{{.FunctionName}}({{.DemoFunctionArgs}})

{{.DemoLogLine}}

{{.DemoRetToBytes}}
}

func (v *{{.DemoTypeName}}) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
`
