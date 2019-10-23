package lession2

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func RollDice() string {

	d1 := randInt(1, 6)
	rand.Seed(time.Now().UTC().UnixNano())
	d2 := randInt(1, 6)
	return strconv.Itoa(d1 + d2)
}

func RollStat(rolls int) {

	s := ""

	for i := 0; i < rolls; i++ {
		s += RollDice() + " "
	}

	for i := 2; i <= 12; i++ {
		num := strconv.Itoa(i)
		count := float64(strings.Count(s, num))

		chance := count / float64(rolls)

		fmt.Printf("\n%d --- %2.1F", i, chance)
	}

}
