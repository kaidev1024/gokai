package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutext sync.RWMutex

func readGo(r <-chan int, idx int) {
	for {
		num := <-r
		fmt.Printf("%dth read number: %d\n", idx, num)
	}

}

func writeGo(w chan<- int, idx int) {
	for {
		num := rand.Intn(1000)
		w <- num
		fmt.Printf("%dth write number: %d\n", idx, num)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGo(ch, i)
	}

	for i := 0; i < 5; i++ {
		go writeGo(ch, i)
	}

	for {
	}
}
