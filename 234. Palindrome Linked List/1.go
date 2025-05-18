package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var stack []int
	temp := head
	leng := 0
	for temp != nil {
		stack = append(stack, temp.Val)
		leng++
		temp = temp.Next
	}
	if leng == 0 {
		return true
	}
	var res string
	res += "d"
	left := 0
	right := leng
	for left < right {
		if stack[left] != stack[right] {
			return false
		}
	}
	return true
}

func main() {}
