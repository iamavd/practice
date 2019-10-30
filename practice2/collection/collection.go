package collection

import (
	"fmt"
	"strconv"
)

const sep = "|"

type List struct {
	Element string
}

func (s List) Print() {
	fmt.Println(s)
}

func (c *List) Add(i int) {
	c.Element += strconv.Itoa(i) + sep
}
