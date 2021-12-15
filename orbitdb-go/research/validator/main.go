package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func call(url string, payload *strings.Reader) (string, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImQxcjEifQ.W5iGCDQh481eTv5MHUmjxwsh5QYTJN9PsTpktFFHtjY")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// fmt.Println(res)
	// fmt.Println(string(body))
	return string(body), nil
}

func main() {

	url := "http://localhost:3000/api/eventlog"
	for i := 0; i < 1000; i++ {
		raw := fmt.Sprintf("{\n\t\"submissionId\": \"%d_f9872d1840D7322E4476C4C08c625Ab9E04d212c1\",\n\t\"signature\": \"f9872d1840D7322E4476C4C08c625Ab9E04d3960f9872d1840D7322E4476C4C08c625Ab9E04d396212\",\n\t\"event\": \"{\\\"signature\\\":\\\"f9872d1840D7322E4476C4C08c625Ab9E04d3960f9872d1840D7322E4476C4C08c625Ab9E04d39601234\\\"}\"\n}", i)
		payload := strings.NewReader(raw)
		res, err := call(url, payload)
		if err != nil {
			fmt.Println(err)
		}
		if i%100 == 0 {
			fmt.Printf("i\t%v\t%v\n%v", i, time.Now().Format("2006.01.02 15:04:05"), res)
		}
	}
}
