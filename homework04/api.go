package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Friends struct {
	Response struct {
		Count int     `json:"count"`
		Users []*User `json:"items"`
	}
}

type User struct {
	UID       int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BDate     string `json:"bdate"`
	Online    int    `json:"online"`
	IsClose   bool   `json:"is_close"`
}

type HistoryMessage struct {
	Response struct {
		Count    int        `json:"count"`
		Messages []*Message `json:"items"`
	}
}

type Message struct {
	Date              int64               `json:"date"`
	FromID            int                 `json:"from_id"`
	ID                int                 `json:"id"`
	Out               int                 `json:"out"`
	PeerID            int                 `json:"peer_id"`
	Text              string              `json:"text"`
	MessageID         int                 `json:"message_id"`
	ForwardedMessages []*ForwardedMessage `json:"fwd_messages"`
	Important         bool                `json:"important"`
	RandomID          int                 `json:"random_id"`
}

type ForwardedMessage struct {
	UID               int                 `json:"user_id"`
	Date              int64               `json:"date"`
	Body              string              `json:"body"`
	ForwardedMessages []*ForwardedMessage `json:"fwd_messages"`
}

func get(url string, maxRetries int, backoffFactor float64) interface{} {
	flag := 0
	for flag < maxRetries {
		resp, err := http.Get(url)
		if err != nil {
			sumTime := backoffFactor * float64(math.Pow(2, float64(flag)))
			flag++
			time.Sleep(time.Duration(sumTime))
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return body
	}
	return nil
}

func getFriends(userID string, fields string) interface{} {
	query := VkConfig["domain"] + "friends.get?" +
		VkConfig["accessToken"] + "&fields=bdate" +
		VkConfig["version"]

	response := get(query, 5, 0.5)

	if response != nil {
		var newResponse = response.([]byte)

		var friendList *Friends
		err := json.Unmarshal(newResponse, &friendList)
		if err != nil {
			log.Fatal("Unmarshal newResponse json failed:", err)
		}
		return friendList.Response.Users
	}
	return nil
}

func getHistoryMessage(userID int64, offset int64, count int64) interface{} {
	query := VkConfig["domain"] + "messages.getHistory?" +
		VkConfig["accessToken"] +
		"&user_id=" + strconv.FormatInt(userID, 10) +
		"&offset=" + strconv.FormatInt(offset, 10) +
		"&count=" + strconv.FormatInt(count, 10) +
		VkConfig["version"]

	response := get(query, 5, 0.5)

	if response != nil {
		var newResponse = response.([]byte)

		var messageList *HistoryMessage
		err := json.Unmarshal(newResponse, &messageList)
		if err != nil {
			log.Fatal("Unmarshal newResponse json failed:", err)
		}
		return messageList.Response.Messages
	}
	return nil
}
