package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func minFromSlice(ss []int) int {
	minItem := ss[0]
	for i := 0; i < len(ss); i++ {
		if ss[i] < minItem {
			minItem = ss[i]
		}
	}
	return minItem
}

/*
2
| \
3  4
| \| \
6  5  7
| \| \| \
4  1  8  3
*/

func minimumTotal(triangle [][]int) int {
	// dp[i, j] 自顶向上走到i,j的最小值
	// dp[i, j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
	count := len(triangle)
	if count == 0 {
		return 0
	}

	dp := make([][]int, count)
	rowCount := len(triangle[count-1])
	for i := 0; i < count; i++ {
		dp[i] = make([]int, rowCount)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < count; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if i == j {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	return minFromSlice(dp[count-1])
}

func main() {
	triangle := [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
