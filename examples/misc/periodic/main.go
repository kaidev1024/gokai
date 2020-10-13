package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var done bool

func main() {
	time.Sleep(1 * time.Second)
	go periodic()
	time.Sleep(3 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()

}

func periodic() {
	fmt.Println("periodic starts...")
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("tick...")
		mu.Lock()
		if done {
			break
		}
		mu.Unlock()
	}
	fmt.Println("periodic ends...")

}
