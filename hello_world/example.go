package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

func main() {

	run_func(quotes, "Quotes")
	run_func(values, "Values")
	run_func(variables, "Variables")
	run_func(constants, "Constants")
	run_func(loops, "For loops")

}

type fn func()

func run_func(f fn, val string) {
	fmt.Println(val)
	f()
	fmt.Println("")
}

func quotes() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}

func values() {

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)

}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func constants() {
	const s string = "constant"

	const n = 5000000

	const d = 3e20 / n

	fmt.Println(s)

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
