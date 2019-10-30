package main

import (
	"fmt"
	"practice2/collection"
)

func main() {

	col := collection.List{}

	for i := 0; i < 10; i++ {
		col.Add(i)
		fmt.Println(col.Last)
	}
	col.Add(12)

	col.Print()
	fmt.Println("")
	fmt.Println(col.Get(1).Value)
	fmt.Println("")

	//	col.Print()
	col.Remove(4)
	col.Print()
	//	col.Print()
	//	col.Add(33)
	//	col.Print()
	//col.Print()
	//col.Remove(4)
	//col.Print()
	//fmt.Println(col.Get(3))
	//fmt.Println(col.First())
	//fmt.Println(col.Last())
	//	fmt.Println(col.Length())

}
