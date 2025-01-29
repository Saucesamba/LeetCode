package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1.Next == nil && l2.Next == nil {
		sum := l1.Val + l2.Val
		res := new(ListNode)
		if sum > 9 {
			perv := int(math.Round(float64(sum / 10)))
			vtor := sum % 10
			sec := new(ListNode)
			sec.Val = perv
			res.Val = vtor
			res.Next = sec
		} else {
			res.Val = sum
		}
		return res
	}

	if l1.Next == nil && l1.Val == 0 {
		return l2
	}
	if l2.Next == nil && l2.Val == 0 {
		return l1
	}

	p := l1
	v := l2

	prev := new(ListNode)
	sum := p.Val + v.Val

	perv := int(math.Round(float64(sum / 10)))
	vtor := sum % 10

	//fmt.Println("sum:", sum, "vtoraya: ", vtor, "pervaya: ", perv)
	res := prev
	prev.Val = vtor

	p = p.Next
	v = v.Next
	if p == nil && v != nil && perv == 0 {
		prev.Next = v
	}
	if p != nil && v == nil && perv == 0 {
		prev.Next = p
	}

	for p != nil && v != nil || perv != 0 {
		sum = perv
		next := new(ListNode)

		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if v != nil {
			sum += v.Val
			v = v.Next
		}
		perv = int(math.Round(float64(sum / 10)))
		vtor = sum % 10
		//fmt.Println("sum:", sum, "vtoraya: ", vtor, "pervaya: ", perv)
		next.Val = vtor
		prev.Next = next
		prev = next
		if p == nil && v != nil && perv == 0 {
			prev.Next = v
			break
		}
		if p != nil && v == nil && perv == 0 {
			prev.Next = p
			break
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	list1, _ := in.ReadString('\n')
	list1 = strings.TrimSpace(list1)
	list2, _ := in.ReadString('\n')
	list2 = strings.TrimSpace(list2)

	mas1 := strings.Split(list1, " ")
	mas2 := strings.Split(list2, " ")

	masInt1 := make([]int, len(mas1))
	for i := 0; i < len(mas1); i++ {
		masInt1[i], _ = strconv.Atoi(mas1[i])
	}
	masInt2 := make([]int, len(mas2))
	for i := 0; i < len(mas2); i++ {
		masInt2[i], _ = strconv.Atoi(mas2[i])
	}

	prev1 := new(ListNode)
	l1 := new(ListNode)
	l1 = prev1

	for i := 0; i < len(mas1); i++ {
		newNode := &ListNode{Val: masInt1[i]}
		if i == 0 {
			l1 = newNode
			prev1 = newNode
		} else {
			prev1.Next = newNode
			prev1 = newNode
		}
	}

	prev2 := new(ListNode)
	l2 := new(ListNode)
	l2 = prev2
	for i := 0; i < len(mas2); i++ {
		newNode := &ListNode{Val: masInt2[i]}
		if i == 0 {
			l2 = newNode
			prev2 = newNode
		} else {
			prev2.Next = newNode
			prev2 = newNode
		}
	}
	output := addTwoNumbers(l1, l2)
	f := output
	for f != nil {
		fmt.Fprintln(out, f.Val)
		f = f.Next
	}
}
