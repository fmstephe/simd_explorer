package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var (
	flagPackage           = flag.String("package", "", "The name of the package")
	flagInstruction       = flag.String("instruction", "", "The assembly name of the instruction to be demonstrated")
	flagSizeClass         = flag.Int("size-class", -1, "The size class of the instruction being demonstrated. Many SIMD instructions work across a range of register sizes.")
	flagDiscriminator     = flag.String("discriminator", "", "A discriminator (can be empty) useful when to demonstrate two versions of an instruction in the same size class, e.g. 'k' ")
	flagTypeDiscriminator = flag.String("type-discriminator", "", "Temporary flag used to add discriminator to type. Will be removed later when unnecessary use of discriminators is removed.")
)

type templateValues struct {
	// Basic Data
	PackageName       string
	InstructionUpper  string
	SizeClass         int
	Discriminator     string
	TypeDiscriminator string

	// Derived Data
	FunctionName      string
	FunctionNameTitle string
	DemoTypeName      string

	// File Names
	AssemblyFileName          string
	StubFileName              string
	AssemblyGeneratorFileName string
	DemoFileName              string
}

func main() {
	flag.Parse()
	validateFlags()

	pkg := strings.ToLower(*flagPackage)
	instructionLower := strings.ToLower(*flagInstruction)
	instructionUpper := strings.ToUpper(*flagInstruction)
	instructionTitle := strings.Title(instructionLower)
	sizeClass := *flagSizeClass
	discriminatorLower := strings.ToLower(*flagDiscriminator)
	discriminatorTitle := strings.Title(discriminatorLower)
	discriminatorUpper := strings.ToUpper(*flagDiscriminator)
	// File names without discriminator unless needed
	var fileNameSuffix string
	if discriminatorLower != "" {
		fileNameSuffix = fmt.Sprintf("%s_%d_%s", instructionLower, sizeClass, discriminatorLower)
	} else {
		fileNameSuffix = fmt.Sprintf("%s_%d", instructionLower, sizeClass)
	}

	tValues := &templateValues{
		PackageName:               pkg,
		InstructionUpper:          instructionUpper,
		SizeClass:                 sizeClass,
		Discriminator:             discriminatorLower,
		DemoTypeName:              fmt.Sprintf("%s%d%s", instructionUpper, sizeClass, typeDiscriminator),
		FunctionName:              fmt.Sprintf("%s%d%s", instructionLower, sizeClass, discriminatorTitle),
		FunctionNameTitle:         fmt.Sprintf("%s%d%s", instructionTitle, sizeClass, discriminatorTitle),
		AssemblyFileName:          fmt.Sprintf("asm_%d_%s_%s.s", sizeClass, discriminatorLower, instructionLower),
		StubFileName:              fmt.Sprintf("stub_%d_%s_%s.go", sizeClass, discriminatorLower, instructionLower),
		AssemblyGeneratorFileName: fmt.Sprintf("asm_%d_%s_%s.go", sizeClass, discriminatorLower, instructionLower),
		DemoFileName:              fmt.Sprintf("%s_%d_%s.go", instructionLower, sizeClass, discriminatorLower),
	}

	fmt.Printf("%#v\n\n", tValues)

	buildDirectories(tValues)
	buildGenerator(tValues)
	buildDemo(tValues)
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
	if *flagDiscriminator == "" {
		fmt.Fprintf(os.Stderr, "Missing -description-short flag value\n")
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

func buildGenerator(tValues *templateValues) {
	fmt.Println("Generator", tValues)

	f, err := os.Create(tValues.PackageName + "/_generate/" + tValues.AssemblyGeneratorFileName)
	if err != nil {
		panic(err)
	}

	generatorTemplate, err := template.New("generator").Parse(asmTemplate)
	if err != nil {
		panic(err)
	}

	err = generatorTemplate.Execute(f, tValues)
	if err != nil {
		panic(err)
	}
}

func buildDemo(tValues *templateValues) {
	fmt.Println("\n\nDemo", tValues)

	f, err := os.Create(tValues.PackageName + "/" + tValues.DemoFileName)
	if err != nil {
		panic(err)
	}

	instructionDemoTemplate, err := template.New("demo").Parse(instructionDemoSource)
	if err != nil {
		panic(err)
	}

	err = instructionDemoTemplate.Execute(f, tValues)
	if err != nil {
		panic(err)
	}
}
