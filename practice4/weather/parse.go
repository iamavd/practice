package weather

import "time"

func encodeTime(utc int64) string {
	return time.Unix(utc, 0).Format("15:04")
}

func printDate(utc int64) string {
	dt := time.Unix(utc, 0)
	switch dt.Day() {
	case time.Now().Day():
		return "Сегодня " + dt.Format("15:04")
	case time.Now().AddDate(0, 0, 1).Day():
		return "Завтра " + dt.Format("15:04")
	default:
		return dt.Format("02.01.2006 15:04")
	}
	return ""
}

func windDirection(deg int) string {
	switch {
	case deg > 45, deg < 135:
		return "восточный"
	case deg > 135, deg < 225:
		return "южный"
	case deg > 225, deg < 315:
		return "западный"
	case deg > 315, deg < 45:
		return "северный"
	}
	return ""
}