// deadlock caused by

// 1. single go routine
// 2. order between channels
// 3. multichannel cross

package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	ch := make(chan int)
	ch <- 23 // deadlock caused by write block
	num := <-ch
	fmt.Println("num: ", num)
}

func main2() {
	ch := make(chan int)

	num := <-ch
	fmt.Println("num: ", num)
	go func() {
		ch <- 23
	}()
}

func main3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}

func printer(str string) {
	mutex.Lock()
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond * 300)
	}
	mutex.Unlock()
}

var ch = make(chan int)

func person1() {
	printer("hello ")
	// ch <- 1
}

func person2() {
	// <-ch
	printer("world")
}

var mutex sync.Mutex

func main() {
	go person2()
	go person1()
	go person2()
	for {
	}
}
