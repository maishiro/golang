package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"

	"github.com/maishiro/golang/httpRequest/store"
)

func main() {
	now := time.Now()
	timestamp := now.Format(time.RFC3339) // 2016-03-25T19:05:54+09:00
	fmt.Println(timestamp)

	url := "http://demo.redmine.org/projects/test/issues.xml"
	respBody := http_get(url)
	fmt.Println(respBody)

	doc, err := xmlquery.Parse(strings.NewReader(respBody))
	// doc, err := xmlquery.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	total_count := xmlquery.FindOne(doc, "/issues/@total_count")
	fmt.Printf("total: %s\n", total_count.InnerText())
	fmt.Printf("\n")

	var db store.Records
	err = db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	issues := xmlquery.Find(doc, "//issue")
	for _, issue := range issues {
		subject := ""
		id := ""
		status_id := ""
		assigned_to_name := ""
		estimated_hours := ""

		if n := issue.SelectElement("//subject"); n != nil {
			subject = n.InnerText()
		}
		if n := issue.SelectElement("//id"); n != nil {
			id = n.InnerText()
		}
		if n := issue.SelectElement("//status/@id"); n != nil {
			status_id = n.InnerText()
		}
		if n := issue.SelectElement("//assigned_to/@name"); n != nil {
			assigned_to_name = n.InnerText()
		}
		if n := issue.SelectElement("//estimated_hours"); n != nil {
			estimated_hours = n.InnerText()
		}

		fmt.Printf("subject %s\n", subject)
		fmt.Printf("id %s\n", id)
		fmt.Printf("status id %s\n", status_id)
		fmt.Printf("assigned_to name %s\n", assigned_to_name)
		fmt.Printf("estimated_hours %s\n", estimated_hours)
		fmt.Printf("\n")

		var issue store.Issue
		issue.TimeStamp = timestamp
		issue.Subject = subject
		issue.Id = id
		issue.StatusId = status_id
		issue.AssignedToName = assigned_to_name
		issue.EstimatedHours = estimated_hours
		db.Add(issue)
	}

}

func http_get(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// req.Header.Set("Accept", "application/xml")
	// 認証情報をセット
	// req.SetBasicAuth("id", "password")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	respBody := string(byteArray)
	return respBody
}
