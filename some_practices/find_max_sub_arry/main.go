package main

import (
	"fmt"
)

func maxSubsequenceSum(a []int) (int, []int) {
	maxSum := 0
	curSum := 0
	curSIndex := 0
	sIndex := 0
	eIndex := 0
	for index, num := range a {
		curSum += num
		if curSum > maxSum {
			maxSum = curSum
			sIndex = curSIndex
			eIndex = index
		} else if curSum <= 0 {
			curSum = 0
			curSIndex = index + 1
		}
	}
	if maxSum == 0 {
		return 0, nil
	}
	return maxSum, a[sIndex : eIndex+1]
}

func main() {
	a1 := []int{-2, -11, 1, -4, -1, 1, -5, 0, -35, 4, 1, -2} // 5
	a2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}               // 最大和的子数组是4, −1, 2, 1 和为 6

	sum1, sub1 := maxSubsequenceSum(a1)
	fmt.Printf("%v max sub is %v sum is %d\n", a1, sub1, sum1)

	sum2, sub2 := maxSubsequenceSum(a2)
	fmt.Printf("%v max sub is %v sum is %d\n", a2, sub2, sum2)
}
