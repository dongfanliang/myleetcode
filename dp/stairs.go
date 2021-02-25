// #15 https://leetcode-cn.com/problems/climbing-stairs/submissions/
package main

import "fmt"

var memo []int

func climbStairs(n int) int {
	// f(n) = f(n-1) + f(n-2)
	memo = make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}

	return work(n)
}

func work(n int) int {
	if n <= 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = work(n-1) + work(n-2)
	return memo[n]
}

func climbStairsDp(n int) int {
	if n < 2 {
		return 1
	}

	x, y := 1, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return y
}

func main() {
	n := 3

	// 递归 + cache
	res := climbStairs(n)
	fmt.Println(res)

	ress := climbStairsDp(n)
	fmt.Println(ress)

}
