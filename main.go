package main 

import (
	"linkedlist/algo"
	"fmt"
	 "strconv"
)

func main() {
	
		var linkedlist algo.LinkedList
		
		for i:=0;i<10;i++{
			node := new(algo.Node)
			node.Setdata(strconv.Itoa(i))
			linkedlist.AddNodeAtTheEnd(node)
		}
		
		node := new(algo.Node)
		node.Setdata(strconv.Itoa(100))
		linkedlist.AddNodeAtTheBegining(node)
		
		node=new(algo.Node)
		node.Setdata(strconv.Itoa(200))
		linkedlist.AddNodeAtIndex(node,10)
		
		
		node=linkedlist.FindNodeWithValue("101")
		
		if(node==nil){
			fmt.Println("Node not found")
		}else{
			fmt.Println("Node found")
		}
		
		fmt.Println(linkedlist.ToString())
		
		
		linkedlist.DeleteNodeWithValue("1000")
		
		
		fmt.Print("After deleting ")
		fmt.Print(linkedlist.ToString())
		
}