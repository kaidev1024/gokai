package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	count := 0
	for i := 0; i < 1000; i++ {
		go func() {
			count += 1
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count: ", count)
}

func main() {
	count := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			count += 1
			mu.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count: ", count)
}
