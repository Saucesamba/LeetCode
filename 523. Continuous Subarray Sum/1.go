package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	mp := map[int]struct{}{}
	mp[0] = struct{}{}
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
		if _, ok := mp[nums[i]%k]; ok {
			return true
		}
		mp[nums[i-1]%k] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{1, 2, 3, 4, 5, 6}, 4))
}
