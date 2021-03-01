package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func onceMaxProfit(prices []int) int {
	// f[i]: 第i天卖出的最大收益
	// curMin : 前i-1天的价格最小值
	// f[i] = max(f[i-1], peices[i]-curMin)
	if len(prices) < 2 {
		return 0
	}

	curMin, res := prices[0], 0
	for i := 1; i < len(prices); i++ {
		res = max(prices[i]-curMin, res)
		curMin = min(prices[i], curMin)
	}
	return res
}

func main() {
	data := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(onceMaxProfit(data))
}
