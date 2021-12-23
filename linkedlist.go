package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type doublylinkedlist struct {
	count int
	head  *Node
	tail  *Node
}

func (l *doublylinkedlist) Append(s string) {

	n := &Node{value: s}

	if l.Size() == 0 {

		l.head = n
		l.tail = n
		l.count = 1
		return
	} else {

		l.head = n
		l.tail.next = n
		return
	}
	l.count++

}

func (l *doublylinkedlist) Size() (count int) {
	c := l.count
	return c
}

func (l *doublylinkedlist) Print() {
	current := l.head
	fmt.Printf("%+v\n", current.value)
	for current.next != nil {
		current = current.next
		fmt.Printf("%+v\n", current.value)
	}
}

func main() {

	k := doublylinkedlist{}
	values := []string{"ser", "frt", "ghy", "yuo", "lop"}

	for _, v := range values {
		k.Append(v)
	}

	k.Print()

	//fmt.Println(k)

}
