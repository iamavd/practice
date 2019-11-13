package bot

import (
	"fmt"
	"os"
)

type English struct {
	Name string
}

func (e English) SayHello() {
	fmt.Println("Yo,it is speaker. Hello here")
}

func (e English) CurrentTime() {
	fmt.Println("Current time")
	printTime("Europe/London", "15:04")
}

func (e English) CurrentDate() {
	fmt.Println("Date - ??????")

}

func (e English) Today() {
	fmt.Println("Today - ??????")
}

func (e English) Bye() {
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

func (e English) Introduce() {
	fmt.Println("Hello, my name is", e.Name)
}
