package main

import (
	"fmt"
	"time"
)

func main1() {
	fmt.Println(time.Now())

	myTimer := time.NewTimer(time.Second * 2)

	nowTime := <-myTimer.C
	fmt.Println(nowTime)
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				quit <- true
				return //runtime.Goexit()
			}
		}
	}()

	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}

	<-quit
	fmt.Println("finish!!!")
}
