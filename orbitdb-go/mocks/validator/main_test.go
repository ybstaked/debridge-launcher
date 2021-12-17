package main

import (
	"fmt"
	"testing"
)

func BenchmarkCall(b *testing.B) {
	url := "http://139.59.164.64:3000/api/eventlog"

	for i := 0; i < b.N; i++ {
		_, err := call(url, i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
