package main

import (
	"fmt"
	"os"
	"practice3/bot"
	"strings"
)

func main() {
	fmt.Println("Hello")
	var language string
	_, err := fmt.Scanln(&language)
	if err != nil {
		fmt.Println("Input error")
		os.Exit(0)
	}

	newBot := createBot(language)

	newBot.Introduce()
	
	for true {
		var reply string
		fmt.Scanln(&reply)
		bot.Listner(newBot, reply)
	}

}

func createBot(language string) bot.Speaker {
	var s bot.Speaker

	switch strings.ToLower(language) {
	case "english", "eng", "en":
		s = bot.English{"John"}
	case "russian", "rus", "ru":
		s = bot.Russian{"Иван"}
	default:
		fmt.Println("OOOPS, we have no such bot =(")
		os.Exit(0)
	}

	return s
}
