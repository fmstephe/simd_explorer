package generate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

func parseAllParams(args string) []*number.Parameter {
	parts := strings.Split(args, ",")

	parameters := []*number.Parameter{}
	for _, part := range parts {
		parameters = append(parameters, parseParam(part))
	}

	return parameters
}

func parseParam(arg string) *number.Parameter {
	arg = strings.TrimSpace(arg)
	parts := strings.Split(arg, " ")
	if len(parts) != 2 {
		panic(fmt.Errorf("bad arg %q, needs two parts separated by a space", arg))
	}

	fieldName := strings.TrimSpace(parts[0])
	typeDec := strings.TrimSpace(parts[1])

	return processTypeDec(fieldName, typeDec)
}

var captureSize = regexp.MustCompile(`^\*\[(\d+)\](.*)$`)

func processTypeDec(fieldName, typeDec string) *number.Parameter {
	if strings.Contains(typeDec, "[") && strings.Contains(typeDec, "]") {
		capturesAll := captureSize.FindAllStringSubmatch(typeDec, -1)
		if len(capturesAll) != 1 {
			panic(fmt.Errorf("Bad array type declaration %q", typeDec))
		}
		captures := capturesAll[0]
		if len(captures) != 3 {
			panic(fmt.Errorf("Bad array type declaration %q", typeDec))
		}
		arrSizeStr := captures[1]
		arrSize, err := strconv.Atoi(arrSizeStr)
		if err != nil {
			panic(err)
		}
		if arrSize <= 1 {
			panic(fmt.Errorf("bad array type declaration, array must be sized 2 or larger: %q", typeDec))
		}
		typeName := captures[2]
		return buildParameter(fieldName, arrSize, typeName)
	} else {
		arrSize := 1
		typeName := typeDec
		return buildParameter(fieldName, arrSize, typeName)
	}
}

func buildParameter(fieldName string, arrSize int, typeName string) *number.Parameter {
	base := chooseBase(fieldName)
	switch typeName {
	case "int8":
		return number.NewNamedIntParameter(fieldName, 8*arrSize, 8, base)
	case "uint8":
		return number.NewNamedUintParameter(fieldName, 8*arrSize, 8, base)
	case "byte":
		return number.NewNamedUintParameter(fieldName, 8*arrSize, 8, base)
	case "int16":
		return number.NewNamedIntParameter(fieldName, 16*arrSize, 16, base)
	case "uint16":
		return number.NewNamedUintParameter(fieldName, 16*arrSize, 16, base)
	case "int32":
		return number.NewNamedIntParameter(fieldName, 32*arrSize, 32, base)
	case "uint32":
		return number.NewNamedUintParameter(fieldName, 32*arrSize, 32, base)
	case "int64":
		return number.NewNamedIntParameter(fieldName, 64*arrSize, 64, base)
	case "uint64":
		return number.NewNamedUintParameter(fieldName, 64*arrSize, 64, base)
	case "float32":
		return number.NewNamedFloatParameter(fieldName, 32*arrSize, 32)
	case "float64":
		return number.NewNamedFloatParameter(fieldName, 64*arrSize, 64)
	default:
		panic(fmt.Errorf("unknown typeName %s", typeName))
	}
}

func chooseBase(fieldName string) int {
	// It looks a bit odd to write this as a switch
	// but I expect this to expand in the future
	switch {
	case strings.HasPrefix(fieldName, "mask"):
		return 16
	case strings.HasPrefix(fieldName, "pred"):
		return 16
	default:
		return 10
	}
}
