package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {

	l, sum := 0, 0
	minLen := len(nums) + 1

	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if minLen > r-l+1 {
				minLen = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	list, _ := in.ReadString('\n')
	list = strings.TrimSpace(list)

	mas := strings.Split(list, " ")
	intMas := make([]int, len(mas))
	for i := 0; i < len(mas); i++ {
		intMas[i], _ = strconv.Atoi(mas[i])
	}
	var targ int
	fmt.Fscanln(in, &targ)
	res := minSubArrayLen(targ, intMas)
	fmt.Fprintln(out, res)
}
