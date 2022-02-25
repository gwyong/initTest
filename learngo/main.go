package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}
func getPage(pageNum int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(pageNum*50)
	// fmt.Println("Requesting: ", pageURL)
	resp, err := http.Get(pageURL)
	checkErr(err)
	checkStatus(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		fmt.Println(id)
		title := card.Find("h2>span").Text()
		fmt.Println(title)
		location := card.Find(".companyLocation").Text()
		fmt.Println(location)
		summary := card.Find(".job-snippet").Text()
		fmt.Println(summary)
	})
}

func getPages(url string) int {
	pages := 0
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkStatus(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = (s.Find("a").Length())
	})
	return pages

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", resp.StatusCode)
	}
}

// func clearnString(txt string) string {

// }
