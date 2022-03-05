package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	location string
	title    string
	summary  string
}

// scrape indeed by a term.
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}
	for i := 0; i < totalPages; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}
	writeJobs(jobs)
	fmt.Println("Done. # extracted:", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF} // for Korean
	file.Write(utf8bom)                 // for Korean

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"ID", "TITLE", "LOCATION", "SUMMARY"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find("h2>span").Text())
	location := CleanString(card.Find(".companyLocation").Text())
	summary := card.Find(".job-snippet").Text()
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		summary:  summary}
}

func getPage(pageNum int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(pageNum*50)
	// fmt.Println("Requesting: ", pageURL)
	resp, err := http.Get(pageURL)
	checkErr(err)
	checkStatus(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func getPages(url string) int {
	pages := 0
	resp, err := http.Get(url)
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

// Cleanstring cleans string
func CleanString(txt string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(txt)), " ")
}
