package main

import "fmt"

func firstUniqChar(s string) int {
	counters := make(map[rune]int)
	for _, x := range s {
		counters[x]++
	}
	fmt.Println(counters)
	for i := 0; i < len(s); i++ {
		if counters[rune(s[i])] == 1 {
			return i
			break
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))

}
