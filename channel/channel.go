package main

import "fmt"

func Fibnocci(n int) {
	if n == 0 {
		fmt.Println("invalid number passed")
		return
	}

	if n == 1 || n == 2 {
		fmt.Println("%d", n)
		return
	}

	prev1 := 1
	prev2 := 0
	for i := 1; i < n; i++ {
		curr := prev1 + prev2
		fmt.Println(curr)

		prev2 = prev1
		prev1 = curr
	}
}

func main() {
	Fibnocci(5)
}
