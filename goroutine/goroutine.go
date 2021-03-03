package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(ii int) { // race condition!(go run -race goroutine.go)
			for {
				fmt.Printf("Hello from goroutine %d\n", ii)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
