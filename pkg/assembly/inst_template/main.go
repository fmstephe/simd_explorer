package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var (
	flagPackage       = flag.String("package", "", "The name of the package")
	flagInstruction   = flag.String("instruction", "", "The assembly name of the instruction to be demonstrated")
	flagSizeClass     = flag.String("size-class", "", "The size class of the instruction being demonstrated. Many SIMD instructions work across a range of register sizes.")
	flagDiscriminator = flag.String("discriminator", "", "A discriminator (can be empty) useful when to demonstrate two versions of an instruction in the same size class, e.g. 'k' ")
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
	FunctionNameCamel string
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

	allInstructions := strings.Split(*flagInstruction, ",")
	allSizeClasses := strings.Split(*flagSizeClass, ",")
	allDiscriminators := strings.Split(*flagDiscriminator, ",")
	if len(allDiscriminators) == 0 {
		allDiscriminators = []string{""}
	}

	for _, instruction := range allInstructions {
		for _, sizeClassStr := range allSizeClasses {
			sizeClass, err := strconv.Atoi(sizeClassStr)
			if err != nil {
				panic(err)
			}
			for _, discriminator := range allDiscriminators {
				buildPermutation(*flagPackage, instruction, discriminator, sizeClass)
			}
		}
	}
}

func buildPermutation(pkg, instruction, disciminator string, sizeClass int) {
	pkg = strings.ToLower(pkg)
	instructionLower := strings.ToLower(instruction)
	instructionUpper := strings.ToUpper(instruction)
	//lint:ignore SA1019 The strings Title function is good enough for our limited purposes
	instructionTitle := strings.Title(instructionLower)
	discriminatorLower := strings.ToLower(disciminator)
	//lint:ignore SA1019 The strings Title function is good enough for our limited purposes
	discriminatorTitle := strings.Title(discriminatorLower)
	discriminatorUpper := strings.ToUpper(disciminator)
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
		DemoTypeName:              fmt.Sprintf("%s%d%s", instructionUpper, sizeClass, discriminatorUpper),
		FunctionName:              fmt.Sprintf("%s%d%s", instructionLower, sizeClass, discriminatorTitle),
		FunctionNameCamel:         fmt.Sprintf("%s%d%s", instructionTitle, sizeClass, discriminatorTitle),
		AssemblyGeneratorFileName: fmt.Sprintf("asm_%s.go", fileNameSuffix),
		AssemblyFileName:          fmt.Sprintf("asm_%s.s", fileNameSuffix),
		StubFileName:              fmt.Sprintf("stub_%s.go", fileNameSuffix),
		DemoFileName:              fmt.Sprintf("demo_%s.go", fileNameSuffix),
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
	if *flagSizeClass == "" {
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
