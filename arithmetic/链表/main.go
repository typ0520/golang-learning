package main

import "fmt"

func main() {
	/**
	  ["addAtHead",
	  "deleteAtIndex",
	  "addAtHead",
	  "addAtHead",
	  "addAtHead",
	  "addAtHead",
	  "addAtHead",
	  "addAtTail",
	  "get",
	  "deleteAtIndex",
	  "deleteAtIndex"]
	  [2],
	  [1],
	  [2],
	  [7],
	  [3],
	  [2],
	  [5],
	  [5],
	  [5],
	  [6],
	  [4]]
	  **/
	obj := Constructor()
	obj.AddAtHead(4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.ValueList())
}
