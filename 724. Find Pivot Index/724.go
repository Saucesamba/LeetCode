package main

func pivotIndex(nums []int) int {
	var pSum []int
	for i, _ := range nums {
		if i == 0 {
			pSum = append(pSum, nums[i])
		} else {
			pSum = append(pSum, pSum[i-1]+nums[i])
		}
	}
	var pivo int = -1
	for i := 0; i < len(pSum); i++ {
		if pSum[i]-nums[i] == pSum[len(pSum)-1]-pSum[i] {
			pivo = i
			break
		}
	}

	if pivo == -1 {
		if pSum[len(pSum)-1]-pSum[0] == 0 {
			pivo = 0
		}
	}
	return pivo
}
