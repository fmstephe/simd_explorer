package main

import (
	"fmt"

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
	return loadsStr
}

func generateAvoInstructionArgs(parameters []*number.Parameter) (loadsStr string) {
	argsStr := ""
	for i, param := range parameters {
		if param.Name() == "ret" {
			argsStr += paramRegister(param)
		} else {
			argsStr += paramRegister(param)
		}
		if i != len(parameters)-1 {
			argsStr += ", "
		}
	}
	return argsStr
}

func generateReturnStore(parameters []*number.Parameter) (returnStore string) {
	for _, param := range parameters {
		if param.Name() == "ret" {
			returnStore = storeToReturn(param)
		}
	}
	return returnStore
}

func generateVZeroUpper(parameters []*number.Parameter) string {
	for _, param := range parameters {
		if param.TotalBitWidth() >= 256 {
			return "\tComment(\"Clear upper halves after YMM usage\")\n\tVZEROUPPER()"
		}
	}
	return ""
}

func paramForLoad(param *number.Parameter) string {
	// e.g. vals1 := Load(Param("vals1"), GP64())

	if param.IsPointer() {
		return fmt.Sprintf("\t%s := Load(Param(%q), GP64())\n", param.Name(), param.Name())
	}

	switch param.TotalBitWidth() {
	case 64:
		return fmt.Sprintf("\t%s := Load(Param(%q), GP64())\n", param.Name(), param.Name())
	case 32:
		return fmt.Sprintf("\t%s := Load(Param(%q), GP32())\n", param.Name(), param.Name())
	case 16:
		return fmt.Sprintf("\t%s := Load(Param(%q), GP16())\n", param.Name(), param.Name())
	case 8:
		return fmt.Sprintf("\t%s := Load(Param(%q), GP8())\n", param.Name(), param.Name())
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

	regLoadStr := fmt.Sprintf("\tComment(\"Load %s into %s register\")\n", param.Name(), registerType)
	regLoadStr += fmt.Sprintf("\t%s := %s()\n", regVarName, registerType)
	regLoadStr += fmt.Sprintf("\tVMOVDQU(Mem{Base: %s}, %s)\n", param.Name(), regVarName)

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
	regLoadStr := fmt.Sprintf("\n\t%s := %s()\n", regVarName, registerType)

	return regLoadStr
}

func paramRegister(param *number.Parameter) string {
	suffix, _ := findRegisterType(param)
	return fmt.Sprintf("%s%s", param.Name(), suffix)
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
	regLoadStr := "\tComment(\"Write results into return memory address\")\n"
	regVarName := fmt.Sprintf("%s%s", param.Name(), suffix)
	regLoadStr += fmt.Sprintf("\tVMOVDQU(%s, Mem{Base: %s})\n", regVarName, param.Name())

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
		return "", ""
	}
}
