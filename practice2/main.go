package main

import (
	"fmt"
	"practice2/collection"
)

func main() {
	col := collection.List{}

	for i := 0; i < 10; i++ {
		col.Add(i)
	}

	col.Print()
	col.Remove(4)
	col.Print()
	fmt.Println(col.Get(3).Value())
	fmt.Println(col.First())
	fmt.Println(col.Last())
	fmt.Println(col.Length())
}
