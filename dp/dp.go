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

func FibnocciDP(n int, dp []int) int {
	if n < 0 {
		fmt.Println("invalid number")
		return -1
	}

	if n == 0 || n == 1 {
		dp[n] = n
		return dp[n]
	}

	if dp[n] >= 0 {
		fmt.Println(dp[n]) 
		return dp[n]
	}

	dp[n] = FibnocciDP(n-2, dp) + FibnocciDP(n-1, dp)
	return dp[n]
}

func main() {
	x := make([]int, 10)
	for i := range x {
		x[i] = -1
	}

	FibnocciDP(9, x)
	//Fibnocci(5)
}
