package main

import (
	"GolangInterfaces/internal/collection"
	"fmt"
)

func main() {
	var array = collection.NewArrayList()
	for i := 0; i < 10; i++ {
		array.Add(i)
	}
	fmt.Printf("%v \n", array)
	array.Reverse()
	fmt.Printf("%v \n", array)
	fmt.Printf("%v \n", array.Contains(5))
	fmt.Printf("%v \n", array.Contains(20))
	var _, ga = array.Get(2)
	fmt.Printf("%v \n", ga)

	var linkedList = collection.LinkedList{}
	for i := 10; i < 20; i++ {
		linkedList.Add(i)
	}
	linkedList.Print()
	linkedList.Reverse()
	linkedList.Print()
	fmt.Printf("%v \n", linkedList.Contains(25))
	fmt.Printf("%v \n", linkedList.Contains(5))
	var _, gl = linkedList.Get(2)
	fmt.Printf("%v \n", gl)
}
