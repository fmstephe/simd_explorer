# An Interactive Reference for SIMD Programming in Go

## To Run

Without cloning the repository run
```
go run github.com/fmstephe/simd_explorer/cmd/simdexplorer@latest
```

OSX (tahoe or later)
```
ROSETTA_ADVERTISE_AVX=1 GOARCH=amd64 go run github.com/fmstephe/simd_explorer/cmd/simdexplorer@latest
```

If you have the source locally run
```
go run cmd/simd_explorer/main.go
```

OSX (tahoe or later)
```
ROSETTA_ADVERTISE_AVX=1 GOARCH=amd64 go run cmd/simdexplorer/main.go
```

## User guide

The simd_explorer app displays a list of SIMD instructions. You can select an instruction to jump to an interactive demonstration of the behaviour of that SIMD instruction.

Type the name of the instruction you are trying to find and the instruction list will be fuzzy filtered. There are over 600 individual demos in the demo list, so scrolling through them one at a time will be very slow.

The registers for any instruction are named parameters to an assembly function which will execute that instruction. The UI builds inputs and outputs so you can change what is in the registers being fed into the demonstrated instruction. All ui inputs/outputs are in either base 2, 10, 16 for integers, or decimal for floating point. The representation is determined by their function in the instruction.

When you open an instruction demo you can return to the instruction list by hitting F1.

When viewing an instruction demo you can view the assembly code being run by the demo by hitting F2 (F1 will return you to the demo screen).

Instruction demo inputs are prefilled by default. You can use F-keys to quickly change the values set (avoiding having to manually edit each field)

    * F3 Zeros every input
    * F4 Autofills every input with an incrementing value (default)
    * F5 Autofills every input with a decrementing value - F4 but reversed

# Supported Architectures

Right now only AMD64 CPUs are supported. The ARM64 NEON instructions are not currently available in Go assembly and the feature detection doesn't appear to work when running emulated AMD64 on M-series macs (which is a shame). The project currently fails to build for ARM64 targets. I would like to have this tool work natively and emulated on OSX, and that may come in the future.

## Skipped Instructions (and rationale)

Some instructions are intentionally omitted because their effects are not easily demonstrable in this UI. The tool focuses on straightforward data manipulation instructions.

- LDMXCSR, STMXCSR
  - Load/store MXCSR (SSE control/status). Change rounding modes, exception masks, DAZ/FTZ, and flags. Just too complex to easily demonstrate.
- MOVNTQ, MOVNTPS, MOVNTDQ, MOVNTDQA
  - Non‑temporal stores. Writing data to memory and hinting that the data should not be kept in the CPU cache. Subtle, hard to demonstrate.
- MASKMOVQ
  - Streaming masked store with write‑combining behaviour; primarily performance/caching behaviour, not a data transform we can display clearly.
- PREFETCH0, PREFETCH1, PREFETCH2, PREFETCHNTA
  - Prefetch hints influence CPU caching. Subtle and hard to demonstrate.
- SFENCE
  - Store fence/ordering primitive. Cross-thread memory ordering instruction. Very subtle and hard to demonstrate.
- CVTPI2PS, CVTTPS2PI
  - MMX↔SSE bridge conversions. Not supported by Go’s assembler (and AVO).

## Misc

Currently AVX/2 is fully covered, other SIMD instructions sets have very spotty coverage.

Alongside the simd_explorer app, there are also generate_avo and regenerate_avo apps. These are used for generating the source code for instruction demos. You won't need to worry about them unless you want to build more instruction demos yourself.

The demos in the project mostly work well - but a large number of them have only been very superficially tested. If you use this tool and find problems or think there is a bug please open an issue, I'll be quite motivated to fix it. One likely source of bugs is a badly chosen size or base/number representation for the inputs and outputs.

The assembly functions used _do_ effectively demonstrate how these SIMD instructions work. But they don't demonstrate how to write a good assembly function in Go. Parameters are mostly passed in as pointers to arrays, this is fine, but the return value is also passed in as a pointer argument, this is unusual. These choices were made to make it easier to generate code for the assembly functions, but likely isn't a very good approach for real world assembly functions. Another note of caution is that we use VMOVDQU/A to load parameters into registers (and also to write return values back to memory) this may not be the most performant approach to doing this.

When I started this project I was trying to teach myself SIMD programming in Go. I found it difficult and very time consuming to identify useful instructions and then to understand their behaviour using the Intel (and Felix Cloutier's website) documentation. So I made this tool to help with this. When I began I genuinely there would be roughly 30-40 instructions to do, and then I would get back to actually using them. I'm not sure how many instructions are here, but there are over 600 demos bundled into this tool (instructions with immediate values get multiple demos of the same instruction and different immediate values). This was a much larger undertaking than I had expected and the biggest challenge ended up being code generation and code style management. But it was fun :)

# Thanks

Egon Elbre (https://github.com/egonelbre)

# TODO

Revisit the way we manage arguments to assembly functions. We may want to use letters A, B... instead of numbers 1, 2... for repeated argument names. We may want to actually return our return values, instead of making them pointer arguments which get set as a side-effect of the function.

We should also revisit how we deal with ZF,CF etc. flags. Right now we combine them into a single integer value - and the result is really hard to interpret.

I am dissatisfied with the current fuzzy filtering of the lists. Find another library for this (avoid the temptation to build your own).
