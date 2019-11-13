package bot

import (
	"fmt"
	"time"
)

func printTime(name string, format string) {
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	fmt.Println(t.Location(), t.Format(format))
}

func date(format string) string {
	return time.Now().Format(format)
}

func localizeDate(t time.Time, monthsVocabulary [12]string) string {
	return fmt.Sprintf("%v %v %v", t.Day(), monthsVocabulary[t.Month()-1], t.Year())
}

func localizeDay(t time.Time, daysVocabulary [7]string) string {
	return daysVocabulary[t.Weekday()]
}

func now() time.Time {
	return time.Now()
}