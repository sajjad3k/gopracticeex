package main

import "fmt"

type node struct {
	data  string
	left  *node
	right *node
}

type binarytree struct {
	root *node
}

//tree
func (t *binarytree) insert(data string) {

	if t.root == nil {
		t.root = &node{data: data}
	} else {
		t.root.insert(data)
	}

}

// node
func (n *node) insert(data string) {

	if n.data <= data {
		if n.left == nil {
			n.left = &node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: data}
		} else {
			n.right.insert(data)
		}
	}

}

func printPreOrder(n *node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%s ", n.data)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%s ", n.data)
	}
}
func main() {

	var t binarytree

	items := []string{"abc", "cde", "efg", "hjk", "yui", "tyu", "qwe"}

	for _, val := range items {

		t.insert(val)

	}

	printPreOrder(t.root)
	fmt.Println("next")
	printPostOrder(t.root)

}
