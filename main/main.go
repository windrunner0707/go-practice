package main

import "awesomeProject/leetcode"

func main() {
	nums := []int{2, 7, 11, 15}
	sum := leetcode.TwoSum(nums, 26)
	println(sum[0])
	println(sum[1])
}
