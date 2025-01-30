package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//надо как-то запоминать позицию буквы, которая повторяется

func lengthOfLongestSubstring(s string) int {
	if s == " " {
		return 1
	}
	maxLen := 0
	resStr := ""
	for _, x := range s {
		if strings.ContainsRune(resStr, x) {
			ind := strings.IndexRune(resStr, x)
			resStr = resStr[ind:]
		} else {
			resStr += string(x)
		}
		if len(resStr) > maxLen {
			maxLen = len(resStr)
		}
		fmt.Println(resStr)
	}
	return maxLen
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var input string
	fmt.Fscan(in, &input)
	res := lengthOfLongestSubstring(input)
	fmt.Fprintln(out, res)
}
