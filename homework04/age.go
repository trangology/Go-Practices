package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getDate(date string) (day, month, year int64) {
	firstIndex := strings.Index(date, ".")
	lastIndex := strings.LastIndex(date, ".")

	day, _ = strconv.ParseInt(date[:firstIndex], 10, 64)
	month, _ = strconv.ParseInt(date[firstIndex+1:lastIndex], 10, 64)
	year, _ = strconv.ParseInt(date[lastIndex+1:], 10, 64)

	return
}

func diff(d1 int, m1 int, y1 int) int {
	today := time.Now()
	var y2, m2, d2 = today.Date()
	years := int(y2 - y1)
	months := int(m2 - time.Month(m1))
	days := int(d2 - d1)

	if days < 0 {
		// days in month:
		t := time.Date(y1, time.Month(m1), 32, 0, 0, 0, 0, time.UTC)
		days += 32 - t.Day()
		months--
	}
	if months < 0 {
		months += 12
		years--
	}
	return years
}

func agePredict(userID string) {
	friends := getFriends(userID, "bdate").([]*User)
	var ageList []int
	var sumAges int

	for _, friend := range friends {
		birthday := friend.BDate

		if len(birthday) >= 8 {
			d, m, y := getDate(birthday)
			age := diff(int(d), int(m), int(y))
			ageList = append(ageList, age)
			sumAges += age
		}
	}
	if len(ageList) != 0 {
		averageAge := sumAges / len(ageList)
		fmt.Print(averageAge)
	}
}

// Add your user ID below
func main() {
	agePredict("")
}
