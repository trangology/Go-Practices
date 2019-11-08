package main

import (
	"net/http"
)

func main() {
	getNews("https://news.ycombinator.com/newest", 5)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/update/", Update)
	http.HandleFunc("/add_label/", AddLabel)
	http.HandleFunc("/recommendations/", Recommendations)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
