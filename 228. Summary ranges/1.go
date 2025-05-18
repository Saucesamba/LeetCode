package main

import "fmt"

func summaryRanges(nums []int) []string {
	left := 0
	right := 0
	var res []string
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 && nums[right] != nums[i] {
			right++
		} else {
			str := fmt.Sprintf("%d->%d", nums[left], nums[right])
			res = append(res, str)
			left = right
			right = i
			if i == len(nums)-1 {
				res = append(res, fmt.Sprintf("%d", nums[i]))
			}
		}
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
