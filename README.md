# A TUI For Learning SIMD Programming in Go

The app found in cmd/simd_explorer/main.go displays a list of SIMD instructions (very incomplete right now). You can select an instruction to jump to an interactive demonstration of the behaviour of that SIMD instruction.

The UI is currently very basic, and likely incomplete from a user's perspective. Right now my focus is on finding the most efficient way to build up the library of assembly functions which demonstrate different SIMD instructions. Once this work is done I will refine the user interface.

## Skipped Instructions (and rationale)

Some instructions are intentionally omitted because their effects are not easily demonstrable in this UI. The tool focuses on straight-forward data manipulation instructions.

- LDMXCSR, STMXCSR
  - Load/store MXCSR (SSE control/status). Change rounding modes, exception masks, DAZ/FTZ, and flags. Just too complex to easily demonstrate.
- MOVNTQ, MOVNTPS
  - Non‑temporal stores. Writing data to memory and hinting that the data should not be kept in the CPU cache. Subtle, hard to demonstrate.
- MASKMOVQ
  - Streaming masked store with write‑combining behaviour; primarily performance/caching behaviour, not a data transform we can display clearly.
- PREFETCH0, PREFETCH1, PREFETCH2, PREFETCHNTA
  - Prefetch hints influence CPU caching. Subtle and hard to demonstrate.
- SFENCE
  - Store fence/ordering primitive. Cross thread memory ordering instruction. Very subtle and hard to demonstrate.
- CVTPI2PS, CVTTPS2PI
  - MMX↔SSE bridge conversions. Not supported by Go’s assembler (and AVO).
