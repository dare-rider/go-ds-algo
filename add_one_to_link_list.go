package main

//Write a function that add 1 to LinkedList number

import (
	"bytes"
	"fmt"
)

type LL struct {
	head *Node
	tail *Node
}

func (l LL) String() string {
	var str bytes.Buffer
	n := l.head
	for n != nil {
		str.WriteString(fmt.Sprintf("%d", n.val))
		n = n.next
	}
	return str.String()
}

type Node struct {
	 val int
	 next *Node
}

func main() {
	n3 := Node{val: 9}
	n2 := Node{9, &n3}
	n1 := Node{9, &n2}
	l := LL{&n1, &n3}
	carry := addOne(l.head, 0)
	if carry != 0 {
		n := Node{carry, l.head}
		l.head = &n
	}
	fmt.Println(l)
}

func addOne(n *Node, level int) int {
	if (n == nil){
		return 1
	}
	carry := addOne(n.next, level + 1)
	n.val += carry
	carry = n.val / 10
	n.val = n.val % 10
	return carry
}