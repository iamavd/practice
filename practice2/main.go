package main

import (
	"fmt"
	. "practice2/collection"
)

func main() {

	col := Collection{
		Element: "aaa",
	}

	col.Element = "d"
	fmt.Println(col.Element)
	col.Print()
}
