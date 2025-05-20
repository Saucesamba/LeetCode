package main

import "strings"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	res := ""
	for _, word := range arr {
		var rewWord string
		for i := len(word) - 1; i >= 0; i-- {
			rewWord += string(word[i])
		}
		res += rewWord + " "
	}
	res = strings.TrimSuffix(res, " ")
	return res
}

func main() {

}
