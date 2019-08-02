package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusPlusFunctional(a, b, c int) int {
	return plus(plus(a, b), c)
}

func main() {
	res := plus(1, 2)

	fmt.Println("1 + 2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = plusPlusFunctional(4, 5, 6)
	fmt.Println("4+5+6=", res)
}
