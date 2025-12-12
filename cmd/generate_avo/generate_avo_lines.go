package main

import (
	"fmt"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

func generateParameterLoads(parameters []*number.Parameter) string {
	loadsStr := ""
	for _, param := range parameters {
		loadsStr += paramForLoad(param)
	}
	return loadsStr
}

func generateRegisterLoads(parameters []*number.Parameter) (loadsStr string) {
	for _, param := range parameters {
		if param.Name() == "ret" {
			loadsStr += returnRegister(param)
		} else {
			loadsStr += paramIntoRegister(param)
		}
	}
	return loadsStr, writeReturn
}

func generateReturnStore(parameters []*number.Parameter) (returnStore string) {
	for _, param := range parameters {
		if param.Name() == "ret" {
			returnStore = storeToReturn(param)
		}
	}
	return returnStore
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

		listStr += paramForList(param)
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

		listStr += paramForList(param)
	}
	return listStr
}

func generateVZeroUpper(parameters []*number.Parameter) string {
	for _, param := range parameters {
		if param.TotalBitWidth() >= 256 {
			return "Comment(\"Clear upper halves after YMM usage\")\nVZEROUPPER()"
		}
	}
	return ""
}

func paramForList(param *number.Parameter) string {
	paramType := param.GoType()
	switch {
	case strings.Contains(paramType, "uint"):
		return fmt.Sprintf("%s: number.NewNamedUintParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "int"):
		return fmt.Sprintf("%s: number.NewNamedIntParameter(%q, %d, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth(), param.Base())
	case strings.Contains(paramType, "float"):
		return fmt.Sprintf("%s: number.NewNamedFloatParameter(%q, %d, %d),\n", param.Name(), param.Name(), param.TotalBitWidth(), param.GetBitWidth())
	default:
		panic(fmt.Errorf("unrecognised parameter type: %s", paramType))
	}
}

func paramForLoad(param *number.Parameter) string {
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

func paramIntoRegister(param *number.Parameter) string {
	if !param.IsPointer() {
		return ""
	}

	suffix, registerType := findRegisterType(param)
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)

	regLoadStr := fmt.Sprintf("Comment(\"Load %s into %s register\")\n", param.Name(), registerType)
	regLoadStr += fmt.Sprintf("%s := %s()\n", regVarName, registerType)
	regLoadStr += fmt.Sprintf("VMOVDQU(Mem{Base: %s}, %s)\n", param.Name(), regVarName)

	return regLoadStr
}

func returnRegister(param *number.Parameter) string {
	if param.Name() != "ret" {
		panic(fmt.Errorf("can't generate return register for param not named 'ret': %s", param.Name()))
	}

	if !param.IsPointer() {
		return ""
	}

	suffix, registerType := findRegisterType(param)
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)
	regLoadStr := fmt.Sprintf("\n%s := %s()\n", regVarName, registerType)

	return regLoadStr
}

func storeToReturn(param *number.Parameter) string {
	if param.Name() != "ret" {
		panic(fmt.Errorf("can't generate store to return for param not named 'ret': %s", param.Name()))
	}

	if !param.IsPointer() {
		return ""
	}

	// Comment("Write results into return memory address")
	// VMOVDQU(retReg, Mem{Base: ret})

	suffix, _ := findRegisterType(param)
	regLoadStr := "Comment(\"Write results into return memory address\")\n"
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)
	regLoadStr += fmt.Sprintf("VMOVDQU(%s, Mem{Base: %s})\n", regVarName, param.Name())

	return regLoadStr
}

func findRegisterType(param *number.Parameter) (suffix, registerName string) {
	switch param.TotalBitWidth() {
	case 512:
		return "Z", "ZMM"
	case 256:
		return "Y", "YMM"
	case 128:
		return "X", "XMM"
	default:
		panic(fmt.Errorf("expect SIMD sized parameter, found %q", param.GoType()))
	}
}
