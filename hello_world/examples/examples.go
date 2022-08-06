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
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "m"
	s[2] = "z"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("Declared in one line", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i*j + i
		}
	}
	fmt.Println("2d:", twoD)
}

func Maps() {
	fmt.Println("Also called dicts in other languages")

	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 20
	m["k3"] = 30

	fmt.Println("map: ", m)
	fmt.Println("map key: ", m["k3"])

	delete(m, "k3")
	fmt.Println("map post-delete: ", m)

	_, prs := m["k3"]
	fmt.Println("Present indicator:", prs)

	n := map[string]int{"Mercury": 1, "Venus": 2, "Earth": 3}
	fmt.Println("Map made in one line:", n)

}

func Ranges() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "berry"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "Golang" {
		fmt.Println(i, c)
	}
}

func Functions() {
	plus := func(a int, b int) int {
		return a + b
	}

	prod := func(a, b int) int {
		return a * b
	}

	plus_result := plus(5, 7)
	fmt.Println("Add two numbers:", plus_result)

	prod_result := prod(5, 7)
	fmt.Println("Product:", prod_result)

	vals := func() (int, int, int) {
		return 3, 10, 15
	}

	a, _, c := vals()
	fmt.Println("Multiple return values:", a, c)

	sum_any := func(nums ...int) {
		fmt.Print(nums, " ")
		total := 0

		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}

	sum_any(4, 8)
	sum_any(6, 15, 8, 3, 32)

	num_slice := []int{1, 2, 3, 4}
	sum_any(num_slice...)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Recursion() {
	fmt.Println("4!:", factorial(4))
	fmt.Println("7!:", factorial(7))

	// required for recursive in nexted functions
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println("Fibbonacci 1:", fib(1))
	fmt.Println("Fibbonacci 3:", fib(3))
	fmt.Println("Fibbonacci 5:", fib(5))
	fmt.Println("Fibbonacci 7:", fib(7))

}
