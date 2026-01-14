package generate

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type templateValues struct {
	// Basic Data
	PackageName       string
	InstructionUpper  string
	SizeClass         int
	Discriminator     string
	TypeDiscriminator string
	Args              string

	// Derived Data
	FunctionName      string
	FunctionNameCamel string
	DemoTypeName      string

	// Generated Avo Code
	AvoLoadArgs        string
	AvoLoadRegisters   string
	AvoInstructionArgs string
	AvoWriteReturn     string
	AvoVZeroUpper      string

	// Generated Demo Code
	DemoFields       string
	DemoConstructor  string
	DemoInputs       string
	DemoName         string
	DemoDescription  string
	DemoInitArrays   string
	DemoFunctionArgs string
	DemoLogLine      string
	DemoRetToBytes   string

	// File Names
	AssemblyFileName          string
	StubFileName              string
	AssemblyGeneratorFileName string
	DemoFileName              string
}

func GenerateDemoFiles(pkg, instruction, discriminator, args, description string, sizeClass int, avoArgs string) {
	pkg = strings.ToLower(pkg)
	instructionLower := strings.ToLower(instruction)
	instructionUpper := strings.ToUpper(instruction)
	//lint:ignore SA1019 The strings Title function is good enough for our limited purposes
	instructionTitle := strings.Title(instructionLower)
	discriminatorLower := strings.ToLower(discriminator)
	//lint:ignore SA1019 The strings Title function is good enough for our limited purposes
	discriminatorTitle := strings.Title(discriminatorLower)
	discriminatorUpper := strings.ToUpper(discriminator)
	// File names without discriminator unless needed
	var fileNameSuffix string
	var demoName string
	if discriminatorLower != "" {
		fileNameSuffix = fmt.Sprintf("%s_%d_%s", instructionLower, sizeClass, discriminatorLower)
		demoName = fmt.Sprintf("%s (%d bit) %s", instructionUpper, sizeClass, discriminatorLower)
	} else {
		fileNameSuffix = fmt.Sprintf("%s_%d", instructionLower, sizeClass)
		demoName = fmt.Sprintf("%s (%d bit)", instructionUpper, sizeClass)
	}

	parameters := parseAllParams(args)

	tValues := &templateValues{
		// Basic Data
		PackageName:      pkg,
		InstructionUpper: instructionUpper,
		SizeClass:        sizeClass,
		Discriminator:    discriminatorLower,
		Args:             args,

		// Derived Data
		FunctionName:      fmt.Sprintf("%s%d%s", instructionLower, sizeClass, discriminatorTitle),
		FunctionNameCamel: fmt.Sprintf("%s%d%s", instructionTitle, sizeClass, discriminatorTitle),
		DemoTypeName:      fmt.Sprintf("%s%d%s", instructionUpper, sizeClass, discriminatorUpper),

		// Generated Avo Code Lines
		AvoLoadArgs:        generateParameterLoads(parameters),
		AvoLoadRegisters:   generateRegisterLoads(parameters),
		AvoInstructionArgs: generateAvoInstructionArgs(parameters, avoArgs),
		AvoWriteReturn:     generateReturnStore(parameters),
		AvoVZeroUpper:      generateVZeroUpper(parameters),

		// Generated Demo Code Lines
		DemoFields:       generateDemoFields(parameters),
		DemoConstructor:  generateDemoConstructor(parameters),
		DemoInputs:       generateDemoInputs(parameters),
		DemoName:         demoName,
		DemoDescription:  fmt.Sprintf("%q", description),
		DemoInitArrays:   generateDemoInitArrays(parameters),
		DemoFunctionArgs: generateDemoFunctionArgs(parameters),
		DemoLogLine:      generateDemoLogLine(instructionUpper, sizeClass, parameters),
		DemoRetToBytes:   generateDemoRetToBytes(parameters),

		// File Names
		AssemblyFileName:          fmt.Sprintf("asm_%s.s", fileNameSuffix),
		StubFileName:              fmt.Sprintf("stub_%s.go", fileNameSuffix),
		AssemblyGeneratorFileName: fmt.Sprintf("asm_%s.go", fileNameSuffix),
		DemoFileName:              fmt.Sprintf("demo_%s.go", fileNameSuffix),
	}

	buildDirectories(tValues)
	buildAvoFile(tValues)
	buildDemoFile(tValues)
}

func buildDirectories(tValues *templateValues) {
	err := os.MkdirAll(tValues.PackageName+"/_generate", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func buildDemoFile(tValues *templateValues) {
	f, err := os.Create(tValues.PackageName + "/" + tValues.DemoFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmplt, err := template.New("demoTemplate").Parse(demoTemplate)
	if err != nil {
		panic(err)
	}

	err = tmplt.Execute(f, tValues)
	if err != nil {
		panic(err)
	}
}

func buildAvoFile(tValues *templateValues) {
	f, err := os.Create(tValues.PackageName + "/_generate/" + tValues.AssemblyGeneratorFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tmplt, err := template.New("avoTemplate").Parse(avoTemplate)
	if err != nil {
		panic(err)
	}

	err = tmplt.Execute(f, tValues)
	if err != nil {
		panic(err)
	}
}
