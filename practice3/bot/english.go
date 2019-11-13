package bot

import (
	"fmt"
	"os"
)

type English struct {
	Name string
}

func (e English) SayHello() {
	fmt.Println("Hello, my name is", e.Name)
}

func (e English) currentTime() {
	fmt.Println("Current time")
	printTime("Europe/London", "15:04")
}

func (e English) currentDate() {
	fmt.Println("Today is", date("January 02 2006"))

}

func (e English) today() {
	fmt.Println("Today is", date("Monday"))
}

func (e English) bye() {
	fmt.Println("Bye")
	os.Exit(1)
}

func (e English) misUnderstanding() {
	fmt.Println("Sorry, I dont understand you")
}

func (e English) initLang() []string {
	commands := []string{"Hello", "Time", "Date", "Today", "Bye"}
	return commands
}