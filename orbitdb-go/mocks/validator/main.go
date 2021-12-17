package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func call(url string, i int) (string, error) {
	raw := fmt.Sprintf("{\n\t\"submissionId\": \"%d_f9872d1840D7322E4476C4C08c625Ab9E04d212c1\",\n\t\"signature\": \"f9872d1840D7322E4476C4C08c625Ab9E04d3960f9872d1840D7322E4476C4C08c625Ab9E04d396212\",\n\t\"event\": \"{\\\"signature\\\":\\\"f9872d1840D7322E4476C4C08c625Ab9E04d3960f9872d1840D7322E4476C4C08c625Ab9E04d39601234\\\"}\"\n}", i)
	payload := strings.NewReader(raw)
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

	return string(body), nil
}

func main() {
	prev := time.Now()
	url := "http://139.59.164.64:3000/api/eventlog"
	fmt.Println("Starting calls to ", url)
	for i := 2002; i <= 3000; i++ {
		res, err := call(url, i)
		if err != nil {
			fmt.Println(err)
		}
		if i%100 == 0 {
			cur := time.Now()
			dur := prev.Sub(cur) * -1
			avg1s := 100 * 1000 / dur.Milliseconds()
			fmt.Println("--------------------------------")
			fmt.Println("--> make 100 calls per", dur.Milliseconds(), "Milliseconds", "total", i)
			fmt.Printf("--> avg of: %v q/s\n", avg1s)
			fmt.Printf("--> res.body: %v\n", res)
			fmt.Println("--------------------------------")

			prev = cur
		}
	}
}
