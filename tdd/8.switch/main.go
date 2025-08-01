package main

import "fmt"

// GetDayName visszaadja a nap nevét sorszám alapján
func GetDayName(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func CategorizeDay(day int) string {
	switch day {
	case 1, 3, 5:
		return "Odd weekday"
	case 2, 4:
		return "Even weekday"
	case 6, 7:
		return "Weekend"
	default:
		return "Invalid day of day number"
	}
}

func main() {
	day := 5
	fmt.Println(GetDayName(day))
	fmt.Println(CategorizeDay(day))
}
