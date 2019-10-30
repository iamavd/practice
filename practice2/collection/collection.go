package collection

import "fmt"

type Collection struct {
	Element string
}

func (s Collection) Print() {
	fmt.Println(s)
}
