package main

import (
	"log"
	"os"
)

func main() {
	mainLog := log.New(os.Stdout, "REST:", log.LstdFlags)
	mainLog.Println("main function starts here.....")
}
