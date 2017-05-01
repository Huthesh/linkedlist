package algo

import (

)

type Node struct {
	data string
	next *Node
}

func (node *Node) Setdata( data string) {
	node.data=data
}

func (node *Node) Setnext( next *Node) {
	node.next=next
}

func (node *Node) Next() *Node{
	return node.next
}

func (node *Node) Getdata() string{
	return node.data
}