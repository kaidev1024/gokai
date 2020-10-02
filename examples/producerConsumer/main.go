package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num: ", num)
	}
}

func main() {
	ch := make(chan int, 5)
	go producer(ch)
	consumer(ch)
}
