package main

import (
	"fmt"
	"net"
)

//create user struct
type Client struct {
	C    chan string
	Name string
	Addr string
}

//create global map to store online clients
var onlineMap map[string]Client

//create global channel to pass client message
var message = make(chan string)

func WriteMsgToClient(client Client, conn net.Conn) {
	// listen to client's channel having message or not
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	netAddr := conn.RemoteAddr().String()
	client := Client{make(chan string), netAddr, netAddr}

	onlineMap[netAddr] = client

	go WriteMsgToClient(client, conn)

	message <- "[" + netAddr + "]" + client.Name + "login"

	for {
	}
}

func Manager() {
	onlineMap = make(map[string]Client)
	for {
		msg := <-message

		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listener.Close()

	go Manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err: ", err)
			return
		}
		go HandlerConnect(conn)
	}
}
