package main

import (
	"fmt"
	"sync"
)

func main1() {
	var a string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a = "hello"
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(a)
}

func main() {
	n := 5
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		// go func() {
		// 	fmt.Println(i)
		// 	wg.Done()
		// }()
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main1() {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func(i) {
			println(i)
			done <- true
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
