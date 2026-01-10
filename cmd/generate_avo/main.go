package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fmstephe/simd_explorer/pkg/generate"
)

var (
	flagPackage       = flag.String("package", "", "The name of the package")
	flagInstruction   = flag.String("instruction", "", "The assembly name of the instruction to be demonstrated")
	flagSizeClass     = flag.Int("size-class", -1, "The size class of the instruction being demonstrated. Many SIMD instructions work across a range of register sizes.")
	flagDiscriminator = flag.String("discriminator", "", "A discriminator (can be empty) useful when to demonstrate two versions of an instruction in the same size class, e.g. 'k' ")
	flagArgs          = flag.String("args", "", "The parameters passed into the generated assembly function")
	flagDescription   = flag.String("description", "", "Used to set the description of each instruction demo")
)

func main() {
	flag.Parse()
	validateFlags()
	generate.GenerateDemoFiles(*flagPackage, *flagInstruction, *flagDiscriminator, *flagArgs, *flagDescription, *flagSizeClass)
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
