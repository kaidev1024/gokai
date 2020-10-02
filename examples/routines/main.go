package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 50; i++ {
		fmt.Println("singing....")
		time.Sleep(1 * time.Second)
	}
}

func dance() {
	for i := 0; i < 50; i++ {
		fmt.Println("dancing....")
		time.Sleep(1 * time.Second)
	}
}

var channel = make(chan int)

func printer(str string) {
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(300 * time.Millisecond)
	}
}
func person1() {
	printer("hello")
	channel <- 8
}

func person2() {
	<-channel
	printer(" world")
}
func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("child process: ", i)
			channel <- i
			fmt.Println("child process: ", i)
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		num := <-channel
		fmt.Println("parent process: ", num)
	}
}
