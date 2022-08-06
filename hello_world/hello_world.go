package main

import (
	"fmt"

	e "example/example/examples"
)

func main() {

	run_func(e.Quotes, "Quotes")
	run_func(e.Values, "Values")
	run_func(e.Variables, "Variables")
	run_func(e.Constants, "Constants")
	run_func(e.Loops, "For loops")
	run_func(e.Ifelse, "If-Else conditions")
	run_func(e.Switches, "Switches")
	run_func(e.Arrays, "Arrays")
	run_func(e.Slices, "Slices")
	run_func(e.Maps, "Maps")
	run_func(e.Ranges, "Ranges")
	run_func(e.Functions, "Functions")
	run_func(e.Recursion, "Recursion")

}

func run_func(f func(), val string) {
	fmt.Println(val)
	f()
	fmt.Println("")
}
