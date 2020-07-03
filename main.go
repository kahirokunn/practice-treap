package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Priority int
	Value    int
	Left     *Node
	Right    *Node
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func NewNode(val int) *Node {
	return &Node{Value: val, Priority: rand.Intn(1000000)}
}

func (t *Tree) Add(newVal int) {
	if t.Root == nil {
		t.Root = NewNode(newVal)
		return
	}

	t.Root = t.Root.Add(newVal)
}

func (n *Node) Add(newVal int) *Node {
	if newVal < n.Value {
		if n.Left == nil {
			n.Left = NewNode(newVal)
		} else {
			n.Left = n.Left.Add(newVal)
		}
		if n.Priority < n.Left.Priority {
			return Rotate2R(n)
		}
		return n
	} else {
		if n.Right == nil {
			n.Right = NewNode(newVal)
		} else {
			n.Right = n.Right.Add(newVal)
		}
		if n.Priority < n.Right.Priority {
			return Rotate2L(n)
		}
		return n
	}
}

func (t *Tree) Contains(target int) bool {
	node := t.Root
	for node != nil {
		if target < node.Value {
			node = node.Left
		} else if target > node.Value {
			node = node.Right
		} else {
			return true
		}
	}
	return false
}

func Rotate2L(n *Node) *Node {
	r := *n.Right
	n.Right = r.Left
	r.Left = n
	return &r
}

func Rotate2R(n *Node) *Node {
	l := *n.Left
	n.Left = l.Right
	l.Right = n
	return &l
}

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	t := NewTree()
	for i := 0; i < 1000000; i++ {
		t.Add(rand.Intn(1000000))
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
