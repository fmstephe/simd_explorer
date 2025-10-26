# A TUI For Learning SIMD Programming in Go

The app found in cmd/simd_explorer/main.go displays a list of SIMD instructions (very incomplete right now). You can select an instruction to jump to an interactive demonstration of the behaviour of that SIMD instruction.

The UI is currently very basic, and likely incomplete from a user's perspective. Right now my focus is on finding the most efficient way to build up the library of assembly functions which demonstrate different SIMD instructions. Once this work is done I will refine the user interface.
