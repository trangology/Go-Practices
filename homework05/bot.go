package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var schedule [][]string
var weekDay = [...]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

func getPage(group string) (content *goquery.Document) {
	url := domain + group + "/raspisanie_zanyatiy_" + group + ".htm"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	content, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}
	return
}

func parseSchedule(doc *goquery.Document, day string) [][]string {
	var timeList, locationList, lessonList []string
	doc.Find("table").Each(func(i int, tableTag *goquery.Selection) {
		idName, _ := tableTag.Attr("id")
		if idName == day {

			// find all tags <td>
			tableTag.Find("td").Each(func(i int, TdTag *goquery.Selection) {

				// find all tags <span> in tag <td>
				TdTag.Find("span").Each(func(i int, spanTag *goquery.Selection) {
					className, _ := spanTag.Parent().Attr("class")
					if className == "time" {
						if spanTag.Text() != "День" {
							timeList = append(timeList, spanTag.Text())
						}
					}
					if className == "room" {
						locationList = append(locationList, spanTag.Text())
					}
				})

				// find all tags <dl> in tag <td>
				TdTag.Find("dl").Each(func(i int, tdTag *goquery.Selection) {
					className, _ := tdTag.Parent().Attr("class")
					if className == "room" {
						locationList = append(locationList, tdTag.Text())
					}
					if className == "lesson" {
						strings.Replace(tdTag.Text(), "\t", "", -1)
						lessonList = append(lessonList, tdTag.Text())
					}
				})
			})
		}
	})
	schedule = append(schedule, timeList, locationList, lessonList)
	return schedule
}

func res(req int) (reply string) {
	if req == 0 {
		reply = "Trang, you have no class on this day :))"
	}

	if 0 < req && req < 7 {
		// req * 3 for timeList, locationList and lessonList
		index := req * 3
		for k := 0; k < len(schedule[index]); k++ {
			reply += schedule[index][k] + "\n" + schedule[index+1][k] + "\n" + schedule[index+2][k] + "\n\n"
		}
	}

	if req == 7 {
		for i := 0; i < 21; i = i + 3 {
			if i != 0 {
				reply += strings.ToUpper(weekDay[i/3]) + "\n\n"
				if len(schedule[i]) == 0 {
					reply += "You have no class on this day" + "\n\n\n\n"
				}
			}
			for k := 0; k < len(schedule[i]); k++ {
				reply += schedule[i][k] + "\n" + schedule[i+1][k] + "\n" + schedule[i+2][k] + "\n\n"
			}
		}

	}
	fmt.Print(len(schedule))
	return
}

// this func will return 0 for Sunday through to 6 for Saturday
// and 7 for all days of the week
func processCommand(command string, group string) string {
	f := getPage(group)
	var today, tomorrow int
	request := 7

	if command == "today" {
		today = int(time.Now().Weekday())
	}

	if command == "tomorrow" {
		tomorrow = int(time.Now().Weekday()) + 1
		if tomorrow == 7 {
			request = 0
		}
	}

	for i := 0; i <= 6; i++ {
		date := strconv.Itoa(i) + "day"
		parseSchedule(f, date)
		if weekDay[i] == command || today == int(i) || tomorrow == int(i) {
			request = i
		}
	}

	if command == "all" {
		request = 7
	}

	return res(request)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(teleToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		message := update.Message.Text

		whiteSpace := strings.Index(message, " ")
		group := message[1:whiteSpace]
		command := message[whiteSpace+1:]

		mess := processCommand(command, group)

		/** switch update.Message.Command() {
		case "start":
			reply = "Hey, buddy! I'm cute bot :3"
		case "i am sad":
			reply = "Me too :("
		**/

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)

		bot.Send(msg)
	}
}
