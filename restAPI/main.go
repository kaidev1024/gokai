package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kaidev1024/gokai/restAPI/data"
)

func main() {
	mainLog := log.New(os.Stdout, "REST:", log.LstdFlags)
	mainLog.Println("main function starts here.....")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		mainLog.Println("main root called")
		_ = data.GetAllProducts()
		mainLog.Printf("%#v\n", r.Body)
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Error reading the data", http.StatusBadRequest)
		}
		mainLog.Println(string(d))
	})
	http.ListenAndServe(":8080", nil)
}
