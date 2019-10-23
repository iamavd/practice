package main

import (
	"practice1/lession1"
	"practice1/lession2"
)

func main() {
	s := lession1.CreateSli(4)
	s = lession1.PullSli(s)
	lession1.PrintSli(s)

	lession2.RollStat(1000)

}
