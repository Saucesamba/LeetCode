package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	var res [][3]int
	var setr [][]int
	sort.Ints(nums)
	if nums[0] > 0 {
		return setr
	}

	for i, x := range nums {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right]+x == 0 {
				subArr := [3]int{nums[i], nums[left], nums[right]}
				f := true
				for _, a := range res {
					if a == subArr {
						f = false
					}
				}
				if f {
					res = append(res, subArr)
				}
				left++
				right--
			} else if nums[left]+nums[right]+x > 0 {
				right--
			} else if nums[left]+nums[right]+x < 0 {
				left++
			}

		}
	}
	for _, x := range res {
		a := []int{x[0], x[1], x[2]}
		setr = append(setr, a)
	}
	return setr
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	list, _ := in.ReadString('\n')
	list = strings.TrimSpace(list)

	mas := strings.Split(list, ",")
	intMas := make([]int, len(mas))
	for i := 0; i < len(mas); i++ {
		intMas[i], _ = strconv.Atoi(mas[i])
	}
	arr := threeSum(intMas)
	fmt.Println(arr)
}
