package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSumClosest(nums []int, target int) int {
	var res int
	sort.Ints(nums)

	for i, x := range nums {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			val := nums[left] + nums[right] + x

			if val > target {
				if target-val < target-res {
					res = val
				}
				right--

			} else if val < target {
				if target-val < target-res {
					res = val
				}
				left++
			}
		}
	}
	return res
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
	var tar int
	fmt.Fscanln(in, &tar)
	arr := threeSumClosest(intMas, tar)
	fmt.Println(arr)
}
