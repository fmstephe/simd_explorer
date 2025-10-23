# Naming Convention For Assembly Types And Methods

Each assembly demonstration package stored here will start with an instruction group. For example the package movhps is an instruction group containing both movhps and vmovhps. These instruction families are chosen based on the instruction grouping used in https://www.felixcloutier.com/x86/.


Each instruction family will have a '\_generate' subdirectory. Here we write the AVO programs which generate our assembly code for each instruction demo. The AVO programs each define their own 'main()' function and are executable independently using `go run name_of_avo_file.go`. The entire directory won't compile, due to the multiple 'main()' function definitions.

Each AVO generator has a go:generate line which looks like
```//go:generate go run asm_vmovhps_64.go -out ../asm_vmovhps_64.s -stubs ../stub_vmovhps_64.go -pkg movhps```

We see here an outline of the naming convention for these files.

## AVO Generator File

The generator file is named 'asm_{instruction_name}_{size_class}.go'. It should be noted that some instructions have variations for the same size class. For example VPBROADCAST has a version with two arguments at each size class, but also can take a third argument 'k', so our naming convention in full is 'asm_{instruction}_{size_class}[_{discriminator}]?.go', where for our example 'discriminator' would be 'k'. The discriminator value should be avoided when it is not required.

## Method Stub File

The stub file is likewise named 'stub_{instruction_name}_{size_class}[_{discriminator}]?.go'. The stub method is named '{instruction_name}{size_class}[{discriminator}]?(...)', where the first letter of Discriminator is capitalised to approximate camel case.

## Generated Assembly File

The generated assembly file is named 'asm_{instruction_name}_{size_class}[_{discriminator}]?.s'. The function name here is the same as the stub function name described above.

## The Instruction Demo Type

Now we have our assembly and stub files arranged, we have an instruction demo type which lives in a file named 'demo_{instruction_name}_{size_class}[_{discriminator}]?.go'. e.g.

```demo_vmovhps_64.go```

This file embeds both the generated assembly and stub file using variable declarations. e.g.

```//go:embed asm_vmovhps_64.s
var assemblyVmovhps64 string

//go:embed stub_vmovhps_64.go
var stubVmovhps64 string```

The naming convention for these variable declarations is 'assembly{instruction_name}{size_class}[{discriminator}]?', and 'stub{instruction_name}{size_class}[{discriminator}]?' respectively. The 'instruction_name' and 'discriminator' names have their first character capitalised to approximate camel case.

We then declare a type named '{instruction_name}{size_class}[{discriminator}]?'. Both 'instruction_name' and 'discriminator' are fully capitalised e.g.

type VMOVHPS64 struct {
}

The type has no fields and is completely stateless. The declared type implements the assembly.Instruction interface using pointer receivers.

## Register Usage Conventions

Most SIMD instructions take two registers and perform some operation on them, storing the result in another register. We typically name these registers regX1, regX2 (for XMM-sized registers, regY1... etc. for wider registers). We prefer to store the results in the first of these two registers where possible. For example

MULPS(regX2, regX1)

is preferred because this will store the result in regX1 and

VMULPS(regX2, regX1, regX1)

is preferred for the same reason.

Some instructions have dramatically different behaviour with different argument orderings. In these cases we prefer to arrange register arguments so that the arithmetic expression when written down reads like x1 * x2 (where * is some arithmetic operator). This ordering is preferred even if the results must be stored in the x2 register. For example

SUBPS(regX2, regX1)

which performs x1 = x1 - x2 is preferred here, and matches both our preferred output register and order of operands. Obviously

VSUBPS(regX2, regX1, regX1)

is preferred when the output register can be specified independently.
