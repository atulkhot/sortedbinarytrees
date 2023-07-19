package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data  string
	left  *Node
	right *Node
}

func build_tree() Node {
	a := Node{data: "A"}
	b := Node{data: "B"}
	c := Node{data: "C"}
	d := Node{data: "D"}
	e := Node{data: "E"}
	f := Node{data: "F"}
	g := Node{data: "G"}
	h := Node{data: "H"}
	i := Node{data: "I"}
	j := Node{data: "J"}

	a.left = &b
	a.right = &c

	b.left = &d
	b.right = &e

	e.left = &g

	c.right = &f

	f.left = &h

	h.left = &i
	h.right = &j

	return a
}

func (node *Node) display_indented(indent string, depth int) string {
	spaces := strings.Repeat(indent, depth)
	result := fmt.Sprintf("%s %s\n", spaces, node.data)
	result_left := ""
	result_right := ""

	if node.left != nil {
		result_left = node.left.display_indented(indent, depth+1)
	}
	if node.right != nil {
		result_right = node.right.display_indented(indent, depth+1)
	}
	return result + result_left + result_right
}

func (node *Node) preorder() string {
	result := fmt.Sprintf("%s", node.data)

	if node.left != nil {
		result = result + " " + node.left.preorder()
	}
	if node.right != nil {
		result = result + " " + node.right.preorder()
	}
	return result
}

func (node *Node) inorder() string {
	result := fmt.Sprintf("%s", node.data)

	if node.left != nil {
		result = node.left.inorder() + " " + result
	}
	if node.right != nil {
		result = result + " " + node.right.inorder()
	}
	return result
}

func (node *Node) postorder() string {
	result := ""

	if node.left != nil {
		result = node.left.postorder()
	}
	if node.right != nil {
		right_result := node.right.postorder()
		if result != "" {
			result = result + " " + right_result
		} else {
			result = right_result
		}
	}
	if result != "" {
		return result + " " + fmt.Sprintf("%s", node.data)
	} else {
		return fmt.Sprintf("%s", node.data)
	}
}

type Queue struct {
	top_sentinel    *Cell
	bottom_sentinel *Cell
}

type Cell struct {
	data *Node
	next *Cell
	prev *Cell
}

func make_queue() Queue {
	top_sentinel := Cell{}
	bottom_sentinel := Cell{next: &top_sentinel, prev: &top_sentinel}

	top_sentinel.next = &bottom_sentinel
	top_sentinel.prev = &bottom_sentinel

	list := Queue{top_sentinel: &top_sentinel, bottom_sentinel: &bottom_sentinel}
	return list
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

func (list *Queue) push_top(v *Node) {
	cell := Cell{data: v}
	list.top_sentinel.add_after(&cell)
}

func (list *Queue) enqueue(v *Node) {
	list.push_top(v)
}

func (list *Queue) is_empty() bool {
	return list.top_sentinel.next == list.bottom_sentinel
}

func (list *Queue) dequeue() *Node {
	if list.is_empty() {
		panic("No item to dequeue - Empty queue")
	}
	elem := list.bottom_sentinel.prev.delete()
	return elem.data
}

func (me *Cell) delete() *Cell {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me
}

func (root *Node) breadth_first() string {
	result := ""
	queue := make_queue()
	queue.enqueue(root)
	for !queue.is_empty() {
		p := queue.dequeue()
		result = result + p.data

		if p.left != nil {
			queue.enqueue(p.left)
		}
		if p.right != nil {
			queue.enqueue(p.right)
		}
		if !queue.is_empty() {
			result += " "
		}
	}
	return result
}

func (root *Node) insert_value(v string) {
	if root.data < v {
		if root.right == nil {
			root.right = &Node{data: v}
		} else {
			root.right.insert_value(v)
		}
	} else if root.data > v {
		if root.left == nil {
			root.left = &Node{data: v}
		} else {
			root.left.insert_value(v)
		}
	}
}

func (root *Node) find_value(v string) *Node {
	if root != nil {
		if root.data < v {
			if root.right == nil {
				return nil
			} else {
				return root.right.find_value(v)
			}
		} else if root.data > v {
			if root.left == nil {
				return nil
			} else {
				return root.left.find_value(v)
			}
		}
	}
	return root
}

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insert_value("I")
	root.insert_value("G")
	root.insert_value("C")
	root.insert_value("E")
	root.insert_value("B")
	root.insert_value("K")
	root.insert_value("S")
	root.insert_value("Q")
	root.insert_value("M")

	// Add F.
	root.insert_value("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.find_value(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
