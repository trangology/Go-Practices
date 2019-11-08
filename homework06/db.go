package main

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type HNews struct {
	ID     uint `gorm:"AUTO_INCREMENT"`
	Title  string
	Author string
	URL    string
	Score  string
	Label  string
}

func Update(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./hnews.db")

	defer db.Close()

	if err != nil {
		panic("Cannot access to database")
	}

	db.CreateTable(&HNews{})

	// creat database using sqlite, can view database by DB.Browser
	for i := range news {
		for _, item := range news[i] {
			hnews := HNews{Title: item.title, Author: item.author, URL: item.url, Score: item.score, Label: "none"}
			db.Create(&hnews)
		}
	}

	http.Redirect(w, r, "/home", http.StatusFound)
}

func Home(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./hnews.db")

	defer db.Close()

	if err != nil {
		panic("Cannot access to database")
	}

	tmpl, err := template.ParseFiles("news_template.html")
	check(err)

	hnews := []HNews{}
	db.Where("Label = ?", "none").Find(&hnews)

	// send database to render in template
	err = tmpl.Execute(w, hnews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddLabel(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.RequestURI())
	check(err)

	q := u.Query()
	label := q["label"][0]
	currentID := q["id"][0]
	db, err := gorm.Open("sqlite3", "./hnews.db")
	defer db.Close()
	if err != nil {
		panic("Cannot access to database")
	}

	db.Model(&HNews{}).Where("ID = ?", currentID).Update("Label", label)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func Recommendations(w http.ResponseWriter, r *http.Request) {
	db, _ := gorm.Open("sqlite3", "./hnews.db")

	defer db.Close()

	hnews := []HNews{}
	db.Where("Label <> ?", "none").Find(&hnews)

	var titleList, labelList []string

	for _, item := range hnews {
		titleList = append(titleList, item.Title)
		labelList = append(labelList, item.Label)
	}

	db.Where("Label = ?", "none").Find(&hnews)
	fitData(titleList, labelList)

	for _, item := range hnews {
		newLabel := predict(item.Title)
		currentID := item.ID
		db.Model(&HNews{}).Where("ID = ?", currentID).Update("Label", newLabel)
	}

	tmpl, err := template.ParseFiles("recommendations_template.html")
	check(err)

	hnews = []HNews{}
	db.Find(&hnews)

	err = tmpl.Execute(w, hnews)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
