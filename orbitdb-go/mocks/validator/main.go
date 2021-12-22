package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func call(url string, i int) (string, error) {
	raw := fmt.Sprintf("{\n\t\"submissionId\": \"test__submissionId_%d\",\n\t\"signature\": \"f9872d1840D7322E4476C4C08c625Ab9E04d3960f9872d1840D7322E4476C4C08c625Ab9E04d396212\",\n\t\"payload\": {\n\t\t\"txHash\": \"test__txHash_%d\",\n\t\t\"submissionId\": \"test__submissionId_%d\",\n\t\t\"chainFrom\": 97,\n\t\t\"chainTo\": 1,\n\t\t\"debridgeId\": \"test__debridgeId\",\n\t\t\"receiverAddr\": \"test__receiverAddr\",\n\t\t\"amount\": \"0.02000000\",\n\t\t\"eventRaw\": \"{\\\"k\\\":\\\"v\\\"}\"\n\t}\n}", i, i, i)
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

const (
	LONDON_ADDRESS    = "139.59.164.64"
	FRANKFURT_ADDRESS = "161.35.31.27"
	LOCALHOST_ADDRESS = "127.0.0.1"
)

func main() {
	prev := time.Now()
	url := fmt.Sprintf("http://%s:3000/api/submission", LOCALHOST_ADDRESS)

	fmt.Println("Starting calls to ", url)
	for i := 1; i <= 1000; i++ {
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
