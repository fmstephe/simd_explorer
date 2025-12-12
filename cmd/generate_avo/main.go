package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var (
	flagPackage       = flag.String("package", "", "The name of the package")
	flagInstruction   = flag.String("instruction", "", "The assembly name of the instruction to be demonstrated")
	flagSizeClass     = flag.Int("size-class", -1, "The size class of the instruction being demonstrated. Many SIMD instructions work across a range of register sizes.")
	flagDiscriminator = flag.String("discriminator", "", "A discriminator (can be empty) useful when to demonstrate two versions of an instruction in the same size class, e.g. 'k' ")
	flagArgs          = flag.String("args", "", "The parameters passed into the generated assembly function")
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

	// Generated Code Lines
	LoadArgsAvo      string
	LoadRegistersAvo string
	WriteReturnAvo   string
	VZeroUpperAvo    string

	// File Names
	AssemblyFileName          string
	StubFileName              string
	AssemblyGeneratorFileName string
	DemoFileName              string
}

func main() {
	flag.Parse()
	validateFlags()
	buildAvoGenerator(*flagPackage, *flagInstruction, *flagDiscriminator, *flagArgs, *flagSizeClass)
}

func buildAvoGenerator(pkg, instruction, discriminator, args string, sizeClass int) {
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
	if discriminatorLower != "" {
		fileNameSuffix = fmt.Sprintf("%s_%d_%s", instructionLower, sizeClass, discriminatorLower)
	} else {
		fileNameSuffix = fmt.Sprintf("%s_%d", instructionLower, sizeClass)
	}

	parameters := parseParameterStr(args)
	loadRegistersAvo := generateRegisterLoads(parameters)
	writeReturnAvo := generateReturnStore(parameters)

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

		// Generated Code Lines
		LoadArgsAvo:      generateParameterLoads(parameters),
		LoadRegistersAvo: loadRegistersAvo,
		WriteReturnAvo:   writeReturnAvo,
		VZeroUpperAvo:    generateVZeroUpper(parameters),

		// File Names
		AssemblyFileName:          fmt.Sprintf("asm_%s.s", fileNameSuffix),
		StubFileName:              fmt.Sprintf("stub_%s.go", fileNameSuffix),
		AssemblyGeneratorFileName: fmt.Sprintf("asm_%s.go", fileNameSuffix),
		DemoFileName:              fmt.Sprintf("demo_%s.go", fileNameSuffix),
	}

	buildDirectories(tValues)
	buildAvoFile(tValues)
}

func validateFlags() {
	if *flagPackage == "" {
		fmt.Fprintf(os.Stderr, "Missing -package flag value\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagInstruction == "" {
		fmt.Fprintf(os.Stderr, "Missing -instruction flag value\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *flagSizeClass == -1 {
		fmt.Fprintf(os.Stderr, "Missing -size-class flag value\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func buildDirectories(tValues *templateValues) {
	err := os.MkdirAll(tValues.PackageName+"/_generate", os.ModePerm)
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

	generatorTemplate, err := template.New("avoTemplate").Parse(avoTemplate)
	if err != nil {
		panic(err)
	}

	err = generatorTemplate.Execute(f, tValues)
	if err != nil {
		panic(err)
	}
}
