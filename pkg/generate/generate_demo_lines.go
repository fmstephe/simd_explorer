package generate

import (
	"fmt"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

func generateDemoFields(parameters []*number.Parameter) string {
	fieldsStr := ""
	for _, param := range parameters {
		fieldsStr += paramForField(param)
	}
	return fieldsStr
}

func generateDemoConstructor(parameters []*number.Parameter) string {
	constructorStr := ""
	for _, param := range parameters {
		constructorStr += paramForConstructor(param)
	}
	return constructorStr
}

func generateDemoInputs(parameters []*number.Parameter) string {
	inputsStr := ""
	for _, param := range parameters {
		if param.Name() == "ret" {
			continue
		}

		inputsStr += paramForList(param)
	}
	return inputsStr
}

func generateDemoInitArrays(parameters []*number.Parameter) string {
	initStr := ""
	for _, param := range parameters {
		initStr += paramForArrayInit(param)
	}
	return initStr
}

func generateDemoFunctionArgs(parameters []*number.Parameter) string {
	argsStr := ""
	for i, param := range parameters {
		if param.IsPointer() {
			argsStr += "&"
		}

		argsStr += param.Name()

		if i != len(parameters)-1 {
			argsStr += ","
		}
	}
	return argsStr
}

func generateDemoLogLine(instructionName string, parameters []*number.Parameter) string {
	logBody := fmt.Sprintf("%s ", instructionName)
	for i, param := range parameters {
		if i == len(parameters)-1 {
			logBody += fmt.Sprintf("%s %%v", param.Name())
		} else {
			logBody += fmt.Sprintf("%s %%v ", param.Name())
		}
	}

	logArgs := ""
	for i, param := range parameters {
		logArgs += param.Name()
		if i != len(parameters)-1 {
			logArgs += ","
		}
	}

	return fmt.Sprintf("\tlog.Printf(\"%s\", %s)", logBody, logArgs)
}

func generateDemoRetToBytes(parameters []*number.Parameter) string {
	for _, param := range parameters {
		if param.Name() == "ret" {
			return paramForRetToBytes(param)
		}
	}
	panic(fmt.Errorf("Could not find 'ret' parameter"))
}

func paramForField(param *number.Parameter) string {
	return fmt.Sprintf("\t%s *number.Parameter\n", param.Name())
}

func paramForConstructor(param *number.Parameter) string {
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		return fmt.Sprintf("\t\t%s: number.NewNamedUintParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "int"):
		return fmt.Sprintf("\t\t%s: number.NewNamedIntParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "float"):
		return fmt.Sprintf("\t\t%s: number.NewNamedFloatParameter(%q, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth())
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}

func paramForList(param *number.Parameter) string {
	return fmt.Sprintf("\t\tv.%s,\n", param.Name())
}

func paramForArrayInit(param *number.Parameter) string {
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := [%d]%s{}\n\tcopy(%s[:], number.ToUint%dSlice(v.%s.FlatData()))\n", param.Name(), param.Parts(), param.PartType(), param.Name(), param.PartBitWidth(), param.Name())
		}
		return fmt.Sprintf("\t%s := number.ToUint%d(v.%s.FlatData())\n", param.Name(), param.PartBitWidth(), param.Name())
	case strings.Contains(paramType, "int"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := [%d]%s{}\n\tcopy(%s[:], number.ToInt%dSlice(v.%s.FlatData()))\n", param.Name(), param.Parts(), param.PartType(), param.Name(), param.PartBitWidth(), param.Name())
		}
		return fmt.Sprintf("\t%s := number.ToInt%d(v.%s.FlatData())\n", param.Name(), param.PartBitWidth(), param.Name())
	case strings.Contains(paramType, "float"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := [%d]%s{}\n\tcopy(%s[:], number.ToFloat%dSlice(v.%s.FlatData()))\n", param.Name(), param.Parts(), param.PartType(), param.Name(), param.PartBitWidth(), param.Name())
		}
		return fmt.Sprintf("\t%s := number.ToFloat%d(v.%s.FlatData())\n", param.Name(), param.PartBitWidth(), param.Name())
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}

func paramForRetToBytes(param *number.Parameter) string {
	name := param.Name()
	bytesVar := fmt.Sprintf("%sBytes", name)
	partBitWidth := param.PartBitWidth()
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := number.Uint%dSliceToBytes(%s[:])\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
		}
		return fmt.Sprintf("\t%s := number.Uint%dToBytes(%s)\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
	case strings.Contains(paramType, "int"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := number.Int%dSliceToBytes(%s[:])\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
		}
		return fmt.Sprintf("\t%s := number.Int%dToBytes(%s)\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
	case strings.Contains(paramType, "float"):
		if param.IsPointer() {
			return fmt.Sprintf("\t%s := number.Float%dSliceToBytes(%s[:])\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
		}
		return fmt.Sprintf("\t%s := number.Float%dToBytes(%s)\n\tv.%s.SetData(%s)", bytesVar, partBitWidth, name, name, bytesVar)
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}
