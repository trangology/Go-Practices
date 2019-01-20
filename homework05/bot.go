package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var schedule [][]string
var weekDay = [...]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}


func getPage(group string) (content *goquery.Document) {
	url := domain + group  + "/raspisanie_zanyatiy_" + group + ".htm"
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	content, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil{
		panic(err)
	}
	return
}

func parseSchedule(doc *goquery.Document, day string) [][]string{
	var timeList, locationList, lesionList []string
	doc.Find("table").Each(func(i int, tableTag *goquery.Selection) {
		idName, _ := tableTag.Attr("id")
		if idName == day {
			tableTag.Find("td").Each(func(i int, TdTag *goquery.Selection) {    	// find all tags <td>
				TdTag.Find("span").Each(func(i int, spanTag *goquery.Selection) {     // find all tags <span> in tag <td>
					className, _ := spanTag.Parent().Attr("class")
					if className == "time" { 		// check class name of tag <span>
						if spanTag.Text() != "День" {
							timeList = append(timeList, spanTag.Text())
						}
					}
					if className == "room" { 			// check class name of tag <span>
						locationList = append(locationList, spanTag.Text())
					}
				})

				TdTag.Find("dl").Each(func(i int, tdTag *goquery.Selection) {   // find all tags <dl> in tag <td>
					className, _ := tdTag.Parent().Attr("class")
					if className == "room" {
						locationList = append(locationList, tdTag.Text())
					}
					if className == "lesson" {
						strings.Replace(tdTag.Text(), "\t", "", -1)
						lesionList = append(lesionList, tdTag.Text())
					}
				})
			})
		}
	})
	schedule = append(schedule, timeList, locationList, lesionList)
	return schedule
}

func res(result int) (reply string) {
	if result == 0{
		reply = "Trang, no class on this day :))"
	}

	if 0 < result && result < 7 {
		index := result * 3			// find exactly day in schedule array
		for k := 0; k < len(schedule[index]); k++ {
			reply += schedule[index][k] + "\n" + schedule[index+1][k] + "\n" + schedule[index+2][k] + "\n\n"
		}
	}

	if result == 7{
		for i := 0; i < 21; i = i + 3 {
			if i != 0 {
				reply += strings.ToUpper(weekDay[i/3]) + "\n\n"	// reply day of week has schedule
			}
			for k := 0; k < len(schedule[i]); k++ {
				reply += schedule[i][k] + "\n" + schedule[i+1][k] + "\n" + schedule[i+2][k] + "\n\n"
			}
		}

	}
	return
}


func processCommand(command string, group string) string {
	f := getPage(group)
	var today, tomorrow int
	needReply := 7

	if command == "today" {
		today = int(time.Now().Weekday())
	}

	if command == "tomorrow" {
		tomorrow = int(time.Now().Weekday()) + 1
		if tomorrow == 7{
			tomorrow = 0
		}
	}

	for i := 0; i <= 6; i++{
		date := strconv.Itoa(i) + "day"
		fmt.Print(parseSchedule(f, date))
		if weekDay[i] == command || today == int(i) || tomorrow == int(i){
			needReply = i
		}
	}

	if command == "all" {
		needReply = 7
	}

	return res(needReply)

}

func main()  {
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
		command := message[whiteSpace + 1:]

		res := processCommand(command, group)

		/** switch update.Message.Command() {
		case "start":
			reply = "Hello. I'm cute bot :3"
		case "i am sad":
			reply = "Me too :("
		**/

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)

		bot.Send(msg)
	}
}
