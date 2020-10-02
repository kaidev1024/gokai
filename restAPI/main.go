package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kaidev1024/gokai/restAPI/handlers"
	protos "github.com/kaidev1024/protobuf/protos/currency"
	"google.golang.org/grpc"
)

func main() {
	mainLog := log.New(os.Stdout, "REST:", log.LstdFlags)
	mainLog.Println("main function starts here.....")

	conn, _ := grpc.Dial("localhost:9092")
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)

	ph := handlers.NewProductHandler(mainLog, cc)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", ph)

	newServer := http.Server{
		Addr:         ":8080",
		Handler:      serverMux,
		ErrorLog:     mainLog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go func() {
		mainLog.Fatal(newServer.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	if sig == os.Interrupt {
		mainLog.Println("program interrupted.....")
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	newServer.Shutdown(ctx)
}
