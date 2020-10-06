package main

import (
	"fmt"
	"net"
	"strings"
	"time"
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

func MakeMsg(client Client, msg string) (buf string) {
	buf = "[" + client.Addr + "]" + client.Name + msg
	return
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	hasData := make(chan bool)

	netAddr := conn.RemoteAddr().String()
	client := Client{make(chan string), netAddr, netAddr}

	onlineMap[netAddr] = client

	go WriteMsgToClient(client, conn)

	message <- MakeMsg(client, "login")

	isQuit := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("client %s drop off\n", client.Name)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err: ", err)
				return
			}
			msg := string(buf[:n-1])

			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				for _, user := range onlineMap {
					userInfo := user.Addr + ": " + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				client.Name = newName
				onlineMap[netAddr] = client
				conn.Write([]byte("rename successful\n"))
			} else {
				message <- MakeMsg(client, msg)
			}
			hasData <- true
		}
	}()

	for {
		select {
		case <-isQuit:
			delete(onlineMap, client.Addr)
			message <- MakeMsg(client, "logout")
			return
		case <-hasData:
			// do nothing, it will reset the timer below
		case <-time.After(time.Second * 10):
			delete(onlineMap, client.Addr)
			message <- MakeMsg(client, "logout")
			return
		}
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
