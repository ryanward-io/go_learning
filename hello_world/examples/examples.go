package examples

import (
	"fmt"
	"math"
	"time"

	"rsc.io/quote"
)

func Quotes() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}

func Values() {

	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)

}

func Variables() {
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

func Constants() {
	const s string = "constant"

	const n = 5000000

	const d = 3e20 / n

	fmt.Println(s)

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

func Loops() {
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

func Ifelse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is evenly divisible by 4")
	}

	if num := 6; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "is positive and has multiple digits")
	}
}

func Switches() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Happy weekend!")
	default:
		fmt.Println("Not long now till the weekend!")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("Good morning!")
	case t < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm boolean")
		case int:
			fmt.Println("I'm an integer")
		default:
			fmt.Println("I'm apparently type", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("String")
}

func Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [4]int{2, 4, 6, 8}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("2D: ", twoD)
}

func Slices() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [4]int{2, 4, 6, 8}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("2D: ", twoD)
}
