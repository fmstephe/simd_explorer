# Naming Convention For Assembly Types And Methods

Each assembly demonstration package stored here will start with an instruction group. For example the package movhps is an instruction group containing both movhps and vmovhps. These instruction families are chosen based on the instruction grouping used in https://www.felixcloutier.com/x86/.


Each instruction family will have a '\_generate' subdirectory. Here we write the AVO programs which generate our assembly code for each instruction demo. The AVO programs each define their own 'main()' function and are executable independently using `go run name_of_avo_file.go`. The entire directory won't compile, due to the multiple 'main()' function definitions.

Each AVO generator has a go:generate line which looks like
```//go:generate go run asm_vmovhps_64.go -out ../asm_vmovhps_64.s -stubs ../stub_vmovhps_64.go -pkg movhps```

We see here an outline of the naming convention for these files.

## AVO Generator File

The generator file is named 'asm_{instruction_name}_{size_class}.go'. It should be noted that some instructions have variations for the same size class. For example VPBROADCAST has a version with two arguments at each size class, but also can take a third argument 'k', so our naming convention in full is 'asm_{instruction}_{size_class}[_{discriminator}]?.go', where for our example 'discriminator' would be 'k'. The discriminator value should be avoided when it is not required.

The arguments, and return value, are always expressed as pointers to the desired value. This is done to make the assembly code look the same for everything, and isn't intended as a good way to write assembly functions generally. Numerical arguments are typically named 'vals' (vector) and 'scalar' (single value) , where there are multiple numerical inputs they are named 'vals1', 'vals2' etc. Mask inputs should be named 'mask', control inputs should be named 'control', predicate inputs should be named 'pred'. The final return value should always be named 'ret'.

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

The struct contains all of the *number.Parameter fields needed to provide arguments to its assembly function. The fields must be declared in the same order as they are passed into the assembly function e.g.

type VMOVHPS64 struct {
	lower *number.Parameter
	upper *number.Parameter
	ret   *number.Parameter
}

The declared type implements the assembly.Instruction interface using pointer receivers.

The types of the parameters (defined in the Inputs() and Output() methods) generally have obvious types. But any input which is used as a mask, predicate, bitfields, or any non-numeric data, should be type uint with base 16 'number.NewUintParameter(fullWidth, partWidth, 16)'. Control inputs, a register packed with indices, should always be base 10.

The values returned by Inputs() and Output() must reference the fields declared in the struct. The methods must always be formatted over multiple lines, no single line methods. The slice returned by the Inputs() method should declare each element on a separate line e.g.

func (v *VMOVHPS64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.lower,
		v.upper,
	}
}

func (v *VMOVHPS64) Output() *number.Parameter {
	return v.ret
}

## Register Usage Conventions

Most SIMD instructions take two registers and perform some operation on them, storing the result in another register. We typically name these registers $(arg1)X, $(arg2)X, where X indicates a 128 bit regiseter, use Y and Z to name wider registers. We prefer to store the results in an explicitly named retX/Y/Z register. For example

MULPS(vals2X, vals1X)

is preferred for pre-AVX instructions which do not allow for an explicit destination registe. Where possible we prefer to store the result in the first argument, vals1X. Similarly

VMULPS(vals2X, vals1X, retX)

is preferred for storing the results in the explicitly named retX.

Please note that we have preserved the register order from MULPS here, choosing vals2X, vals1X, even though the order here does not determine which register the output is directed to. This is preferred in order to make all assembly functions as identical as possible, with the only differences being meaningful differences between the instructions demoed. Accidental differences are deliberately minimised.

Some instructions have dramatically different behaviour with different argument orderings. In these cases we prefer to arrange register arguments so that the arithmetic expression when written down reads like x1 * x2 (where * is some arithmetic operator). This ordering is preferred even if the results must be stored in the vasl2X register. For example

SUBPS(vals2X, vals1X)

which performs x1 = x1 - x2 is preferred here, and matches both our preferred output register and order of operands. Obviously

VSUBPS(vals2, vals1X, retX)

is preferred when the output register can be specified independently.

# VZEROUPPER

After completing any 256 or 512 bit operations we should _always_ invoke VZEROUPPER. This is required to break possible dependencies with the upper parts of the YMM registers. This is a behavioural oddity with Intel CPUs treatment of uncleared YMM registers. The performance penalty of failing to do this is _severe_ when 128 bit SSE instructions are executed after 256/512 bit instructions. Any assembly function which uses 256/512 bit SIMD operations should invoke VZEROUPPER before returning. We should likely _never_ mix 128 and 256/512 bit operations in a single function. It's also good to have this as an implementation note for people learning SIMD assembly for the first time. This needs an explanatory comment, and should look like

// YMM/ZMM processing complete, clear upper half of YMM registers
VZEROUPPER

For some clarification around the need and historical background of this please see this discussion [thread](https://community.intel.com/t5/Intel-ISA-Extensions/What-is-the-status-of-VZEROUPPER-use/td-p/1098375)
