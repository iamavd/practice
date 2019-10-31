package collection

import (
	"fmt"
	"strconv"
)

const sep = " "

type List struct {
	first  *Element
	last   *Element
}

type Element struct {
	value string
	next  *Element
	prev  *Element
}

func (c *List) Print() {
	for elem:=c.first; elem.next != nil; elem=elem.next { 
		fmt.Print(elem.next.value, sep)
	}
	fmt.Println("")
}

func (c *List) Add(i int) {
	e := &Element{
		value: strconv.Itoa(i),
	}

	if c.first == nil {
		c.first = e
		c.last = e
	} else {
		last := c.last
		last.next = e
		c.last = e
		c.last.prev = last
	}
}

func (c *List) Get(i int) *Element {
	res := c.first
	for a := 0; a < i; a++ {
		res = res.next
	}
	return res
}

func (c *List) Remove(index int) {
	var elem *Element
	elem = c.Get(index)

	if elem.prev != nil {
		elem.prev.next = elem.next
	} else{
		c.first = elem.next
	}
	elem = nil
}

func (c *List) Length() int {
	i:=0
	for elem:=c.first; elem != nil; elem = elem.next {
		i++
	}
	return i
}

func (c* List) First() string {
	return c.first.value
}

func (c* List) Last() string {
	return c.last.value
}

func (e* Element) Value() string {
	return string(e.value)
}