package algo

import (
"bytes"
)

type LinkedList struct {
	head *Node
}

func (linkedlist *LinkedList) AddNodeAtTheEnd(node *Node) {
	if(node==nil){
		panic("Can not insert nil to linked list")
	}
	if(linkedlist.head==nil){
		linkedlist.head=node
	}else{
		temp:=linkedlist.head
		//Find the last Node, i.e node.next is nil
		for temp.Next()!=nil{
			temp=temp.Next()
		}
		temp.Setnext(node)
	}
}


func (linkedlist *LinkedList) AddNodeAtTheBegining(node *Node) {
	if(linkedlist.head==nil){
		linkedlist.head=node
	}else{
		node.Setnext(linkedlist.head)
		linkedlist.head=node
	}
}

func (linkedlist *LinkedList) AddNodeAtIndex(node *Node,index int) {
	
	if(index<0){
		panic("Index can not be -ve")
	}
	
	if(index==0){
		linkedlist.AddNodeAtTheBegining(node)	
		return
	}
	
	prev:=linkedlist.head
	current:=linkedlist.head
	for i:=0;i<index;i++{
		if(current==nil){
			panic("Index out of bound")
		}
		prev=current
		current=current.Next()
	}
	prev.Setnext(node)
	node.Setnext(current)
}

func (linkedlist *LinkedList) DeleteNodeWithValue(value string) bool{
	
	if(linkedlist.head!=nil && linkedlist.head.Getdata()==value){
		temp:=linkedlist.head
		linkedlist.head=temp.Next()
		temp.Setnext(nil)
		return true
	}
	
	prev:=linkedlist.head
	current:=linkedlist.head
	for current!=nil{
		if(current.Getdata()==value){
			prev.Setnext(current.Next())
			current.Setnext(nil)
			return true
		}
		prev=current
		current=current.Next()
	}
	return false	
}

func (linkedlist *LinkedList) FindNodeWithValue(value string) *Node {
	
	temp:=linkedlist.head	
	for temp!=nil{		
		if(temp.data==value){
			return temp
		}	
		temp=temp.Next()
	}	
	return nil
}

func (linkedlist *LinkedList) ToString() string {
	var buffer bytes.Buffer
	temp:=linkedlist.head
    for  temp!=nil {
        buffer.WriteString(temp.Getdata())
        buffer.WriteString(" ")
        temp=temp.Next();
    }
    return buffer.String()
}