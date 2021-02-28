package utils

import (
	"fmt"
	"time"
)

const (
	Days    = 3
	Hours   = 2
	Minutes = 1
	Seconds = 0
)

// Returns a correct datetime string in Russian
func FormatDateTime(input time.Time) string {
	month := ""
	switch input.Month() {
	case time.January:
		month = "января"
		break
	case time.February:
		month = "февраля"
		break
	case time.March:
		month = "марта"
		break
	case time.April:
		month = "апреля"
		break
	case time.May:
		month = "мая"
		break
	case time.June:
		month = "июня"
		break
	case time.July:
		month = "июля"
		break
	case time.August:
		month = "августа"
		break
	case time.September:
		month = "сентября"
		break
	case time.October:
		month = "октября"
		break
	case time.November:
		month = "ноября"
		break
	case time.December:
		month = "декабря"
		break
	}

	return fmt.Sprintf("%d:%d:%d %d %s %d", input.Hour(), input.Minute(), input.Second(), input.Year(), month, input.Year())
}

// Returns a correct duration in Russian
func FormatDuration(input time.Duration) string {
	s := int(input.Milliseconds() / 1000)
	h := s / 3600
	s -= h * 3600
	m := s / 60
	s -= m * 60

	str := ""
	if h > 0 {
		str += FormatUnit(h, Hours) + " "
	}
	if m > 0 {
		str += FormatUnit(m, Minutes) + " "
	}
	str += FormatUnit(s, Seconds)

	return str
}

// Returns a correct time unit string in Russian
func FormatUnit(num int, unit int) string {
	/*
		0: 1, но не 11: день
		1: 2-4, но не 12-14: дня
		2: остальное + 11-14: дней
	*/
	caseType := 0

	if num%10 == 1 && num%100 != 11 {
		caseType = 0
	} else if num%10 >= 2 && num%10 <= 4 && !(num%100 >= 12 && num%100 <= 14) {
		caseType = 1
	} else {
		caseType = 2
	}

	str := ""
	if caseType == 0 {
		switch unit {
		case Days:
			str = "день"
			break
		case Hours:
			str = "час"
			break
		case Minutes:
			str = "минута"
			break
		case Seconds:
			str = "секунда"
			break
		}
	} else if caseType == 1 {
		switch unit {
		case Days:
			str = "дня"
			break
		case Hours:
			str = "часа"
			break
		case Minutes:
			str = "минуты"
			break
		case Seconds:
			str = "секунды"
			break
		}
	} else {
		switch unit {
		case Days:
			str = "дней"
			break
		case Hours:
			str = "часов"
			break
		case Minutes:
			str = "минут"
			break
		case Seconds:
			str = "секунд"
			break
		}
	}

	return fmt.Sprintf("%d %s", num, str)
}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
