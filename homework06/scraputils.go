package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type News struct {
	title string
	author string
	url string
	score string
}

var news [][]News


func getContent(url string) (body *goquery.Document) {
	res, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	body, err = goquery.NewDocumentFromReader(res.Body)

	if err != nil{
		panic(err)
	}
	return 
}

func extractNews(parser *goquery.Document) []News {
	var n News
	var newsList []News
	var titleList, urlList, scoreList, authorList []string

	parser.Find(".storylink").Each(func(i int, title *goquery.Selection) {
		titleList = append(titleList, title.Text())
		val := title.Nodes[0].Attr[0].Val
		urlList = append(urlList, val)
	})

	parser.Find(".hnuser").Each(func(i int, author *goquery.Selection) {
		authorList = append(authorList, author.Text())
	})

	parser.Find(".score").Each(func(i int, score *goquery.Selection) {
		scoreList = append(scoreList, score.Text())
	})

	for i := 0; i < len(titleList); i++{
		n.title = titleList[i]
		n.author = authorList[i]
		n.url = urlList[i]
		n.score = scoreList[i]
		newsList = append(newsList, n)
	}

	return newsList
}

func extractNewPage(parser *goquery.Document) string{
	link := parser.Find(".morelink")
	result := link.Nodes[0].Attr[0].Val
	return result
}


func getNews(url string, nPages int)  {
	// Collect news from a given web page
	fmt.Print("Collecting data from page: " + url)
	for i := 1; i <= nPages; i++{
		content := getContent(url)
		newsList := extractNews(content)
		nextPage := extractNewPage(content)
		url = "https://news.ycombinator.com/" + nextPage
		news = append(news, newsList)
	}
}
