package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Find(val int) bool {
	cur := ll.Head
	for cur != nil {
		if cur.Val == val {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (ll *LinkedList) Append(node *Node) {
	if ll.Head == nil {
		ll.Head = node
	} else {
		cur := ll.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

func (ll *LinkedList) InsertLeft(leftOf int, node *Node) {
	if ll.Head != nil {
		if ll.Head.Val == leftOf {
			node.Next = ll.Head
			ll.Head = node
		} else {
			cur := ll.Head
			for cur.Next != nil {
				if cur.Next.Val == leftOf {
					node.Next = cur.Next
					cur.Next = node
					break
				}
				cur = cur.Next
			}
		}
	}
}

func (ll *LinkedList) InsertRight(rightOf int, node *Node) {
	if ll.Head != nil {
		cur := ll.Head
		for cur != nil {
			if cur.Val == rightOf {
				node.Next = cur.Next
				cur.Next = node
				break
			}
			cur = cur.Next
		}
	}
}

func (ll *LinkedList) Delete(val int) {
	if ll.Head != nil {
		if ll.Head.Val == val {
			ll.Head = ll.Head.Next
		} else {
			cur := ll.Head
			var prev *Node
			for cur != nil {
				if cur.Val == val {
					prev.Next = cur.Next
					break
				}
				prev = cur
				cur = cur.Next
			}
		}
	}
}

func (ll *LinkedList) Reverse() {
	cur := ll.Head
	var prev *Node
	for cur != nil {
		next := cur.Next // temporary holder for next value
		cur.Next = prev  // flip the pointer to the previous value
		prev = cur       // Keep track of the current pointer for the next iteration
		cur = next       // iterate forward
	}
	ll.Head = prev
}

func (ll *LinkedList) Print() {
	cur := ll.Head
	for cur != nil {
		fmt.Printf("%d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Println("")
}

func main() {
	linkedList := LinkedList{Head: &Node{Val: 10}}
	linkedList.Append(&Node{Val: 11})
	linkedList.Print()
	linkedList.Delete(10)
	linkedList.Print()
	linkedList.InsertLeft(11, &Node{Val: 10})
	linkedList.InsertRight(11, &Node{Val: 12})
	linkedList.InsertLeft(10, &Node{Val: 9})
	linkedList.InsertRight(9, &Node{Val: 10})
	linkedList.Delete(10)
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
}
