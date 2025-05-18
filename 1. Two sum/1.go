package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mMap := make(map[int]int)
	for i, x := range nums {
		comp := target - x
		if ind, ok := mMap[comp]; ok {
			return []int{ind, i}
		}
		mMap[x] = i
	}
	return nil
}
func main() {

	fmt.Println(twoSum([]int{3, 3}, 6))
}
