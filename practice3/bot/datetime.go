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
	fmt.Println(localDate(t))
}

func date(format string) string {
	t := time.Now()
	return t.Format(format)
}

func localDate(t time.Time) string {
	return fmt.Sprintf("%s %02d. %s",
		days[t.Weekday()][:3], t.Day(), months[t.Month()-1][:3],
	)
}
