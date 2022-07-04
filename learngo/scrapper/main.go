package main

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
	title       string
	companyName string
	location    string
	summary     string
}

func (e extractedJob) String() string {
	return "제목 : " + e.title + "\n" + "회사명 :" + e.companyName + "\n" + "위치 :" + e.location + "\n" + "설명 :" + e.summary
}

var baseURL string = "https://kr.indeed.com/jobs?q=golang&l=%EA%B2%BD%EA%B8%B0%EB%8F%84%20%EC%88%98%EC%9B%90&vjk=5ea9ea023c6ccc03"

func main() {

	mainC := make(chan []extractedJob)
	var jobs []extractedJob
	totalPage := getPages()
	fmt.Println(totalPage)

	for i := 0; i < totalPage; i++ {
		go getPage(i, mainC)
	}
	for i := 0; i < totalPage; i++ {
		extratedjobs := <-mainC
		jobs = append(jobs, extratedjobs...)

	}
	fmt.Println(jobs)
	writeJobs(jobs)

}
func getPage(page int, mainC chan<- []extractedJob) {

	c := make(chan extractedJob)
	var eJobs = []extractedJob{}
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchMainCard := doc.Find(".job_seen_beacon")
	searchMainCard.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)

	})

	for i := 0; i < searchMainCard.Length(); i++ {
		job := <-c
		eJobs = append(eJobs, job)
	}

	/*searchShelfContainer := doc.Find(".jobCardShelfContainer")
	searchShelfContainer.Each(func(i int, s *goquery.Selection) {
		eJobs[i].summary = cleanString(s.Find(".job-snippet").Text())
	})*/

	mainC <- eJobs
}

func extractJob(s *goquery.Selection, c chan<- extractedJob) {
	job := extractedJob{
		title:       cleanString(s.Find(".singleLineTitle").Text()),
		companyName: cleanString(s.Find(".companyName").Text()),
		location:    cleanString(s.Find(".companyLocation").Text()),
		summary:     cleanString(s.Find(".job-snippet").Text()),
	}
	c <- job
}
func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Title", "CompanyName", "Location", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.title, job.companyName, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPages() int {

	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()

	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
 	}
}
