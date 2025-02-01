package main

type NumArray struct {
	num  []int
	pSum []int
}

func Constructor(nums []int) NumArray {
	var arr NumArray
	for i := 0; i < len(nums); i++ {
		arr.num = append(arr.num, nums[i])
		if i == 0 {
			arr.pSum = append(arr.pSum, nums[i])
		} else {
			arr.pSum = append(arr.pSum, arr.pSum[i-1]+nums[i])
		}
	}
	return arr
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.pSum[right]
	}
	return this.pSum[right] - this.pSum[left-1]
}
