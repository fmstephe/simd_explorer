package main

import (
	"fmt"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

func generateInputsList(parameters []*number.Parameter) string {
	listStr := ""
	for _, param := range parameters {
		if param.Name() == "ret" {
			continue
		}

		listStr += paramForList(param)
	}
	return listStr
}

func generateRetList(parameters []*number.Parameter) string {
	listStr := ""
	for _, param := range parameters {
		if param.Name() != "ret" {
			continue
		}

		listStr += paramForList(param)
	}
	return listStr
}

func paramForList(param *number.Parameter) string {
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		return fmt.Sprintf("\t%s: number.NewNamedUintParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "int"):
		return fmt.Sprintf("\t%s: number.NewNamedIntParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "float"):
		return fmt.Sprintf("\t%s: number.NewNamedFloatParameter(%q, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth())
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}
