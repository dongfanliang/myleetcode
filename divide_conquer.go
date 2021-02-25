// #50 https://leetcode-cn.com/problems/powx-n/
package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x, n/2)
}
