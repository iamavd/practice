package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Skill int
}

type Match struct {
	player1 Player
	player2 Player
}

func main() {
	start()
}

func start() {
	ch := make(chan int)
	p1 := Player{"Vasya", 23}
	p2 := Player{"Ivan", 44}
	m := Match{p1, p2}

	go m.play(ch)

	for {
		power, ok := <-ch
		if ok {
			fmt.Println("Сила удара:", power)
		} else {
			fmt.Println("Матч окончен")
			break
		}
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func (m Match) play(c chan int) {
	hitNum := 0
	for {
		power := randInt(0, 50)
		c <- power
		hitNum++
		fmt.Println("Номер подачи:", hitNum)
		if hitNum%2 == 0 {
			if m.player1.Skill < power {
				fmt.Println(m.player2.Name, "побеждает")
				break
			} else {
				fmt.Println(m.player1.Name, "отбивает удар")
			}
		} else {
			if m.player2.Skill < power {
				fmt.Println(m.player1.Name, "побеждает")
				break
			} else {
				fmt.Println(m.player2.Name, "отбивает удар")

			}
		}
	}
	close(c)
}
