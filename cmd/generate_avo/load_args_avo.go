package main

import (
	"fmt"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

func generateParameterLoads(parameters []*number.Parameter) string {
	loadsStr := ""
	for _, param := range parameters {
		loadsStr += generateParamForLoad(param)
	}
	return loadsStr
}

func generateRegisterLoads(parameters []*number.Parameter) (loadsStr, writeReturn string) {
	for _, param := range parameters {
		if param.Name() == "ret" {
			loadsStr += generateReturnRegister(param)
			writeReturn = generateStoreToReturn(param)
		} else {
			loadsStr += generateParamIntoRegister(param)
		}
	}
	return loadsStr, writeReturn
}

func generateInputsList(parameters []*number.Parameter) string {
	listStr := ""
	for _, param := range parameters {
		if param.Name() == "ret" {
			// TODO
			// This is a pretty imperfect test, as we simply check here  that the return parameter is named 'ret'
			// If we ever decide that this naming convention is too restrictive this code will break
			continue
		}

		listStr += generateParamForList(param)
	}
	return listStr
}

func generateRetList(parameters []*number.Parameter) string {
	listStr := ""
	for _, param := range parameters {
		if param.Name() != "ret" {
			// TODO
			// Inverse of imperfect check angiushed over in comments above
			continue
		}

		listStr += generateParamForList(param)
	}
	return listStr
}

func generateParamForList(param *number.Parameter) string {
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		return fmt.Sprintf("%s: number.NewNamedUintParameter(%q, %d, %d, %d),\n", param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "int"):
		return fmt.Sprintf("%s: number.NewNamedIntParameter(%q, %d, %d, %d),\n", param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "float"):
		return fmt.Sprintf("%s: number.NewNamedFloatParameter(%q, %d, %d, %d),\n", param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}

func generateParamForLoad(param *number.Parameter) string {
	// e.g. vals1 := Load(Param("vals1"), GP64())

	if param.IsPointer() {
		return fmt.Sprintf("%s := Load(Param(%q), GP64())\n", param.Name(), param.Name())
	}

	switch param.TotalBitWidth() {
	case 64:
		return fmt.Sprintf("%s := Load(Param(%q), GP64())\n", param.Name(), param.Name())
	case 32:
		return fmt.Sprintf("%s := Load(Param(%q), GP32())\n", param.Name(), param.Name())
	case 16:
		return fmt.Sprintf("%s := Load(Param(%q), GP16())\n", param.Name(), param.Name())
	case 8:
		return fmt.Sprintf("%s := Load(Param(%q), GP8())\n", param.Name(), param.Name())
	default:
		panic(fmt.Errorf("unrecognised bit width (%d) for non-pointer parameter %s", param.TotalBitWidth(), param))
	}
}

func generateParamIntoRegister(param *number.Parameter) string {
	if !param.IsPointer() {
		return ""
	}

	suffix, registerType := generateRegisterType(param)
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)

	regLoadStr := fmt.Sprintf("Comment(\"Load %s into %s register\")\n", param.Name(), registerType)
	regLoadStr += fmt.Sprintf("%s := %s()\n", regVarName, registerType)
	regLoadStr += fmt.Sprintf("VMOVDQA(Mem{Base: %s}, %s)\n", param.Name(), regVarName)

	return regLoadStr
}

func generateReturnRegister(param *number.Parameter) string {
	if param.Name() != "ret" {
		panic(fmt.Errorf("can't generate return register for param not named 'ret': %s", param.Name()))
	}

	if !param.IsPointer() {
		return ""
	}

	suffix, registerType := generateRegisterType(param)
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)
	regLoadStr := fmt.Sprintf("\n%s := %s()\n", regVarName, registerType)

	return regLoadStr
}

func generateStoreToReturn(param *number.Parameter) string {
	if param.Name() != "ret" {
		panic(fmt.Errorf("can't generate store to return for param not named 'ret': %s", param.Name()))
	}

	if !param.IsPointer() {
		return ""
	}

	// Comment("Write results into return memory address")
	// VMOVDQA(retReg, Mem{Base: ret})

	suffix, _ := generateRegisterType(param)
	regLoadStr := "Comment(\"Write results into return memory address\")\n"
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)
	regLoadStr += fmt.Sprintf("VMOVDQA(%s, Mem{Base: %s})\n", regVarName, param.Name())

	return regLoadStr
}

func generateRegisterType(param *number.Parameter) (suffix, registerName string) {
	switch param.TotalBitWidth() {
	case 512:
		return "Z", "ZMM"
	case 256:
		return "Y", "YMM"
	case 128:
		return "X", "XMM"
	default:
		return "", ""
	}
}
