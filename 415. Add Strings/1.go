package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	var res []byte
	first := len(num1) - 1
	second := len(num2) - 1
	carry := 0

	for first >= 0 || second >= 0 {
		if first >= 0 {
			carry += int(num1[first] - '0')
			first--
		}
		if second >= 0 {
			carry += int(num2[second] - '0')
			second--
		}
		res = append(res, byte(carry%10)+'0')
		carry /= 10
	}
	if carry > 0 {
		res = append(res, byte(carry%10)+'0')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	fmt.Println(addStrings("11", "123"))
}
