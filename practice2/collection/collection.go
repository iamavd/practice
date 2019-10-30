package collection

import (
	"fmt"
	"strconv"
)

type List struct {
	Length int
	First  *Element
	Last   *Element
}

type Element struct {
	Value string
	Next  *Element
	Prev  *Element
}

func (c *List) Print() {
	res := c.First
	for i := 0; i < 10; i++ { //FIXME
		fmt.Print(res.Next.Value, " ")
		res = res.Next
	}
}

func (c *List) Add(i int) {
	e := &Element{
		Value: strconv.Itoa(i),
	}

	if c.First == nil {
		c.First = e
		c.Last = e
	} else {
		last := c.Last
		last.Next = e
		c.Last = e
		c.Last.Prev = last
	}
	c.Length++
}

func (c *List) Get(i int) *Element {
	res := c.First
	for a := 0; a < i; a++ {
		res = res.Next
	}
	return res
}

func (c *List) Remove(i int) {
	c.Get(i).Value = "1111" //.Value = "-" //c.Get(i).Prev.Prev
}
