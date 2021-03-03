package main

// 二分查找

import (
	"fmt"
)

func maxFromSlice(arr []int) int {
	maxItem := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxItem {
			maxItem = arr[i]
		}
	}
	return maxItem
}

func wroker(arr []int, speed int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		// 向上取整
		sum += (arr[i] + speed - 1) / speed
	}

	return sum
}

func minEatingSpeed(piles []int, H int) int {
	n := len(piles)

	// 速度范围
	minSpeed, maxSpeed := 1, maxFromSlice(piles[:n])
	for minSpeed < maxSpeed {
		speed := (minSpeed + maxSpeed) / 2
		if wroker(piles[:n], speed) <= H {
			maxSpeed = speed
		} else {
			minSpeed = (speed + 1)
		}

	}
	return minSpeed
}

func main() {
	arr := []int{3, 6, 7, 11}
	fmt.Println(minEatingSpeed(arr, 8))
}
