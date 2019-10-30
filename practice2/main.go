package main

import (
	"practice2/collection"
)

func main() {

	col := collection.List{}

	for i := 0; i < 10; i++ {
		col.Add(i)
	}

	col.Print()
}
