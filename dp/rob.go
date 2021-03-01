package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func rob(nums []int) int {
	// sum[i] = max(sum[i-1], sum[i-2]+nums[i])
	count := len(nums)
	if count == 0 {
		return 0
	}
	sumRes := make([]int, count+1)

	if count == 1 {
		return nums[0]
	}

	sumRes[0], sumRes[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < count; i++ {
		sumRes[i] = max(sumRes[i-1], sumRes[i-2]+nums[i])
	}

	fmt.Println(sumRes)
	return sumRes[count-1]
}

func rob2(nums []int) int {
	// sum[i] = max(sum[i-1], sum[i-2]+nums[i])
	count := len(nums)
	if count == 0 {
		return 0
	}

	if count == 1 {
		return nums[0]
	}

	prev, cur := nums[0], max(nums[0], nums[1])
	for i := 2; i < count; i++ {
		fmt.Println(prev, cur)
		prev, cur = cur, max(prev+nums[i], cur)
	}

	return cur
}

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob2(nums))
}
