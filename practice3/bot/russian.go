package bot

import (
	"fmt"
	"os"
)

var days = [...]string{
	"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"}

var months = [...]string{
	"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь",
	"Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
}

type Russian struct {
	Name string
}

func (r Russian) SayHello() {
	fmt.Println("O, privet. Davai rush B")
}

func (r Russian) CurrentTime() {
	//	name := "Europe/Minsk" // Europe/London
	fmt.Println("Текущее время")
	printTime("Europe/Minsk", "15:04")
}

func (r Russian) CurrentDate() {
	fmt.Println("Дата - ??????")
}

func (r Russian) Today() {
	fmt.Println("Сегодня", date("Monday"))
	printTime("Europe/Moscow", "Monday")

}

func (r Russian) Bye() {
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

func (e Russian) Introduce() {
	fmt.Println("Привет, меня зовут", e.Name)
}