package bot

import (
	"fmt"
	"os"
)

var days = [...]string{
	"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"}

var months = [...]string{
	"января", "февраля", "марта", "апреля", "мая", "июня",
	"июля", "августа", "сентября", "октября", "ноября", "декабря",
}

type Russian struct {
	Name string
}

func (r Russian) SayHello() {
	fmt.Println("Привет, меня зовут", r.Name)
}

func (r Russian) currentTime() {
	fmt.Println("Текущее время")
	printTime("Europe/Minsk", "15:04")
}

func (r Russian) currentDate() {
	fmt.Println("Сегодня", localizeDate(now(),months))
}

func (r Russian) today() {
	fmt.Println("Сегодня", localizeDay(now(),days))
}

func (r Russian) bye() {
	fmt.Println("Пока")
	os.Exit(1)
}

func (r Russian) misUnderstanding() {
	fmt.Println("Прости, я тебя не понимаю")
}

func (e Russian) initLang() []string {
	commands := []string{"Привет", "Время", "Дата", "День", "Пока"}
	return commands
}